package docker

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	dockerStatusExited   = "exited"
	dockerStatusRunning  = "running"
	dockerStatusStarting = "starting"
)

type Docker struct {
	ContainerID   string
	ContainerName string
	ContainerStartTimeout int64
}

type MappedPort struct {
	InternalPort int
	ExposedPort  int
}

type ContainerOption struct {
	Name              string
	ContainerFileName string
	Options           map[string]string
	MountVolumePath   string
	MappedPorts []MappedPort
}

func (d *Docker) isInstalled() bool {

	command := exec.Command("docker", "ps")
	err := command.Run()
	if err != nil {

		return false
	}

	return true
}

func (d *Docker) Start(c ContainerOption) (string, error) {

	dockerArgs := d.getDockerRunOptions(c)

	command := exec.Command("docker", dockerArgs...)
	command.Stderr = os.Stderr
	result, err := command.Output()
	if err != nil {
		return "", err
	}
	d.ContainerID = strings.TrimSpace(string(result))
	d.ContainerName = c.Name
	command = exec.Command("docker", "inspect", d.ContainerID)
	result, err = command.Output()
	if err != nil {
		d.Stop()
		return "", err
	}

	time.Sleep(time.Second * time.Duration(d.ContainerStartTimeout))

	return string(result), nil
}

func (d *Docker) WaitForStartOrKill(timeout int) error {

	for tick := 0; tick < timeout; tick++ {

		containerStatus := d.getContainerStatus()
		if containerStatus == dockerStatusRunning {
			return nil
		}
		if containerStatus == dockerStatusExited {
			return nil
		}
		time.Sleep(time.Second)
	}

	d.Stop()

	return errors.New("Docker faile to start in given time period so stopped")
}

func (d *Docker) getContainerStatus() string {

	command := exec.Command("docker", "ps", "-a", "--format", "{{.ID}}|{{.Status}}|{{.Ports}}|{{.Names}}")
	output, err := command.CombinedOutput()

	if err != nil {

		d.Stop()
		return dockerStatusExited

	}

	outputString := string(output)
	outputString = strings.TrimSpace(outputString)
	dockerPsResponse := strings.Split(outputString, "\n")

	for _, response := range dockerPsResponse {

		containerStatusData := strings.Split(response, "|")
		containerStatus := containerStatusData[1]
		containerName := containerStatusData[3]
		if containerName == d.ContainerName {
			if strings.HasPrefix(containerStatus, "Up ") {
				return dockerStatusRunning
			}
		}
	}

	return dockerStatusStarting
}

func (d *Docker) getDockerRunOptions(c ContainerOption) []string {

	var exposedPorts []string

	for _, k := range c.MappedPorts {

		exposedPorts = append(exposedPorts,"-p")
		exposedPorts = append(exposedPorts,fmt.Sprintf("%d:%d",k.ExposedPort,k.InternalPort))
	}

	//portExpose := c.PortExpose //fmt.Sprintf("%s:%s", c.PortExpose, c.PortExpose)

	//var args []string

	for key, value := range c.Options {

		exposedPorts = append(exposedPorts, []string{"-e", fmt.Sprintf("%s=%s", key, value)}...)
	}

	exposedPorts = append(exposedPorts, []string{"--tmpfs", c.MountVolumePath, c.ContainerFileName}...)
	dockerArgs := append([]string{"run", "-d", "--name", c.Name}, exposedPorts... )

	return dockerArgs
}

func (d *Docker) Stop() {

	exec.Command("docker", "rm", "-f", d.ContainerID).Run()
}
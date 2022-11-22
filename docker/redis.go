package docker

import (
	"fmt"
	"time"
)

func (m *Container) StartRedisDocker(port int, pass string) {

	envVar := map[string]string{
		"REDIS_PASSWORD": pass,
	}

	mappedPorts := MappedPort{
		InternalPort: 6379,
		ExposedPort:  port,
	}

	containerOption := ContainerOption{
		Name:              fmt.Sprintf("test-redis-%d",time.Now().Unix()),
		Options:           envVar,
		MountVolumePath:   "/var/lib/redis",
		ContainerFileName: m.ImageName,
		MappedPorts: []MappedPort{mappedPorts},
	}

	m.Docker = Docker{}
	m.Docker.Start(containerOption)
	m.Docker.WaitForStartOrKill(int(m.Docker.ContainerStartTimeout))
}

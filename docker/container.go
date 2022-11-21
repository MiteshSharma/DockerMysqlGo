package docker

const ContainerStartTimeout = 10

type Container struct {
	ImageName   string
	Docker      Docker
}

func (m *Container) Stop() {

	m.Docker.Stop()
}

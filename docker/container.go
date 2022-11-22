package docker

type Container struct {
	ImageName   string
	Docker      Docker
}

func (m *Container) Stop() {

	m.Docker.Stop()
}

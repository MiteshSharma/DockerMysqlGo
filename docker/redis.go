package docker

func (m *Container) StartRedisDocker(port int, pass string) {

	envVar := map[string]string{
		"REDIS_PASSWORD": pass,
	}

	mappedPorts := MappedPort{
		InternalPort: 6379,
		ExposedPort:  port,
	}

	containerOption := ContainerOption{
		Name:              "project-redis-1",
		Options:           envVar,
		MountVolumePath:   "/var/lib/redis",
		ContainerFileName: m.ImageName,
		MappedPorts: []MappedPort{mappedPorts},
	}

	m.Docker = Docker{}
	m.Docker.Start(containerOption)
	m.Docker.WaitForStartOrKill(ContainerStartTimeout)
}

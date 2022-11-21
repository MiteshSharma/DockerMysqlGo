package docker

func (m *Container) StartMysqlDocker(user, password string, port int, dbname string) {

	mysqlOptions := map[string]string{
		"MYSQL_ROOT_PASSWORD": password,
		"MYSQL_USER":          user,
		"MYSQL_PASSWORD":      password,
		"MYSQL_DATABASE":      dbname,
//		"MYSQL_TCP_PORT":      fmt.Sprintf("%d", port),
	}

	mappedPorts := MappedPort{
		InternalPort: 3306,
		ExposedPort:  port,
	}

	containerOption := ContainerOption{
		Name:              "project-mysql-1",
		Options:           mysqlOptions,
		MountVolumePath:   "/var/lib/mysql",
		ContainerFileName: m.ImageName,
		MappedPorts: []MappedPort{mappedPorts},
	}

	m.Docker = Docker{}
	m.Docker.Start(containerOption)
	m.Docker.WaitForStartOrKill(ContainerStartTimeout)
}

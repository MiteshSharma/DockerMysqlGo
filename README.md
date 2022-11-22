# DockerMysqlGo

Running mysql docker container using go code

### Steps to run mysql docker container using command line:

Step 1: Check docker is installed by running command "docker ps". If docker is not installed, install docker from https://docs.docker.com/install/.

Step 2: Start docker mysql container with user name "gouser", password "gopassword" and database name "godb" using command:

```
docker run --name our-mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=gouser -e MYSQL_PASSWORD=gopassword -e MYSQL_DATABASE=godb -p 3306:3306 --tmpfs /var/lib/mysql mysql:5.7
```

Step 3: Check docker status using command with container name our-mysql-container: 

```
docker ps -a
```

Step 4: Once docker is up and running, connect using connection string:

```
gouser:gopassword@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local
```

### Docker using go

We are using same steps to run docker container using go.

docker.go : This file contains generic code to run any type of container image using docker.

mysql.go : This file runs mysql container using docker.go

redis.go : This file runs redis container using docker.go

### environment variables
This library allows setting different environment variables required to run any image, this is passed as ContainerOptions Object

### Exposing Internal Container ports

This library allows exposing multiple internal ports in the image, example you may have already MySQL running on the host machine on port 3306 and you want to spin a test MySQL on another port say 13306, in docker terminal this will look like
```shell script
docker run -p 13306:3306
```
this library allows easily exposing internal ports and mapping them to custom external ports via MappedPort Object

### Run any Image
This library allows running any image, pass any number of environment variables, mount volume and expose multiple ports, with this you should be able to run test on any image

### Example Redis Test Container

Create an afunction that extends `Container` Object then define the requirements for your image

```shell script
go get -u github.com/mudphilo/go-docker@latest
```

```go
import (
	"fmt"
	"github.com/go-redis/redis"
	docker "github.com/mudphilo/go-docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func StartRedisDocker() {

	port := 26973 // custom port
	pass := "pass" // custom password
	imageName := "bitnami/redis"  // redis image name

	envVar := map[string]string{
		"REDIS_PASSWORD": pass, // we set redis password via environment variables
	}

    // lest do port mapping, this will enable us avoid port conflicts with host machine
	mappedPorts := docker.MappedPort{
		InternalPort: 6379,  // we want to expose default redis port to a custom port
		ExposedPort:  port,
	}

    // create your container options
	containerOption := docker.ContainerOption{
		Name:              "project-redis-1",  // container name
		Options:           envVar,
		MountVolumePath:   "/var/lib/redis", // mount volume
		ContainerFileName: imageName,
		MappedPorts: []MappedPort{mappedPorts},
	}
     
    // 
	m.Docker = docker.Docker{}
	m.Docker.Start(containerOption)
	m.Docker.WaitForStartOrKill(ContainerStartTimeout)
    defer m.Docker.Stop()  // when done call this to destroy the container
   
    // go ahead with your testing
  
    uri := fmt.Sprintf("%s:%d", "127.0.0.1", port)
    
    	redisConfig := redis.Options{
    		MinIdleConns: 10,
    		IdleTimeout:  60 * time.Second,
    		PoolSize:     1000,
    		Addr:         uri,
    	}
    
    	redisConfig.Password = pass
    
    	client := redis.NewClient(&redisConfig)
    
    	testKey := "test_key_name"
    	testData := "test data here"
    
    	_, err := client.Set(testKey,testData,time.Minute * 5).Result()
    	assert.NoError(t,err)
    
    	res, err := client.Get(testKey).Result()
    	assert.NoError(t,err)
    
    	assert.Equal(t,testData,res)
}
```
 
 Hurrah you have done your unit tests
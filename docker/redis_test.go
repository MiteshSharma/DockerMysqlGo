package docker

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisContainer(t *testing.T) {

	port := 26973
	pass := "pass"
	imageName := "bitnami/redis"

	docker := Docker{
		ContainerID:   "redis-unittest",
		ContainerName: "redis-unitest",
	}

	redisContainer := Container{
		Docker: docker,
		ImageName: imageName,
	}

	redisContainer.StartRedisDocker(port,pass)
	defer redisContainer.Stop()

	time.Sleep(time.Second * 15)

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
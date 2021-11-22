package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

type Redis struct {
	client *redis.Client
}

func InitRedis() (*Redis, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		redisPort = 6379
	}

	redisUri := fmt.Sprintf("%s:%d", redisHost, redisPort)
	client := redis.NewClient(&redis.Options{
		Addr: redisUri,
		Password: "",
		DB: 0,  //use default DB
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	fmt.Println("Redis init was completed")
	return &Redis{
		client: client,
	}, nil
}


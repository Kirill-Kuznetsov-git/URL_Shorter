package db

import (
	"URLShortener/app/config"
	"context"
	"errors"
	"fmt"
	RedisLibrary "github.com/go-redis/redis"
	"URLShortener/app/hasher"
	"log"
	"os"
	"strconv"
)

type Redis struct {
	client *RedisLibrary.Client
}

func InitRedis() (*Redis, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		redisPort = 6379
	}

	redisUri := fmt.Sprintf("%s:%d", redisHost, redisPort)
	client := RedisLibrary.NewClient(&RedisLibrary.Options{
		Addr: redisUri,
		Password: "",
		DB: 0,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	fmt.Println("Redis init was completed")
	return &Redis{
		client: client,
	}, nil
}

func (redis *Redis) Close() error {
	err := redis.client.Close()
	if err != nil {
		return err
	}
	return nil
}


func (redis *Redis) Save(url ShortUrl) (string, error) {
	url_string := config.StructToString(url)

	res, err := redis.client.Get(url_string).Result()
	if err == RedisLibrary.Nil {
		log.Println("Redis: url will be created", url_string, res)

		err = redis.client.Set(url_string, "0", 0).Err()
		if err != nil {
			return "Redis: set error. Please retry", err
		}
		return hasher.Encode(url.Id), nil
	} else if err != nil {
		return "Redis: get error. Please retry", err
	}
	return "Redis: url already exists", errors.New("already exists")
}

func (redis *Redis) Get(ctx context.Context, URLstring string) (string, error){
	res, err := redis.client.Get(URLstring).Result()
	if err == RedisLibrary.Nil{
		return "Redis: Such url does not exists", errors.New("not exists")
	}
	return res, nil

}
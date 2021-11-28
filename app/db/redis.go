package db

import (
	"URLShortener/app/hasher"
	"context"
	"errors"
	"fmt"
	RedisLibrary "github.com/go-redis/redis/v8"
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
	if _, err := client.Ping(context.Background()).Result(); err != nil {
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


func (redis *Redis) Save(ctx context.Context, UrlOrigin string) (string, error) {
	res, err := redis.client.Get(ctx, UrlOrigin).Result()
	if err == RedisLibrary.Nil {
		log.Println("Redis: url will be created", UrlOrigin, res)
		UrlShort, err := hasher.Encode()
		if err != nil{
			return "Error with hasher", err
		}
		_, err = redis.client.Get(ctx, UrlShort).Result()

		for err != RedisLibrary.Nil{
			UrlShort, _ = hasher.Encode()
			_, err = redis.client.Get(ctx, UrlShort).Result()
		}

		err = redis.client.Set(ctx, UrlShort, UrlOrigin, 0).Err()
		if err != nil{
			return "Redis error", err
		}
		err = redis.client.Set(ctx, UrlOrigin, UrlShort, 0).Err()
		if err != nil{
			return "Redis error", err
		}

		return UrlShort, nil
	} else if err != nil {
		return "Redis: get error. Please retry", err
	}
	return res, nil
}


func (redis *Redis) Get(ctx context.Context, UrlShort string) (string, error){
	log.Println("UrlShort: " + UrlShort)
	UrlOrigin, err := redis.client.Get(ctx, UrlShort).Result()
	if err == RedisLibrary.Nil{
		return "Redis: Such url does not exists", errors.New("not exists")
	}
	if err != nil {
		return "redis error", err
	}
	return UrlOrigin, nil
}
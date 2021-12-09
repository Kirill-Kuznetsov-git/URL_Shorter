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

func (redis *Redis) Init() error {
	// Try to get all requied information for connecting to Redis
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
		return err
	}
	log.Println("Redis init was completed")
	redis.client = client
	return nil
}

func (redis *Redis) Close() error {
	err := redis.client.Close()
	if err != nil {
		return err
	}
	return nil
}


func (redis *Redis) Save(ctx context.Context, UrlOrigin string) (*Url, error) {
	// Check: Maybe such UrlOrigin already exist in DB
	UrlShort, err := redis.client.Get(ctx, UrlOrigin).Result()
	// err == RedisLibrary.Nil means that such recording does noe exist in DB
	if err == RedisLibrary.Nil {
		log.Println("Redis: url will be created", UrlOrigin)

		// Сreate a new hash string until it becomes unique
		UrlShort, _ := hasher.Encode()
		_, err = redis.client.Get(ctx, UrlShort).Result()
		for err != RedisLibrary.Nil{
			UrlShort, _ = hasher.Encode()
			_, err = redis.client.Get(ctx, UrlShort).Result()
		}

		// Create two recording in DB UrlOrigin -> UrlShort and UrlShort -> UrlOrigin
		err = redis.client.Set(ctx, UrlShort, UrlOrigin, 0).Err()
		if err != nil{
			return nil, err
		}
		err = redis.client.Set(ctx, UrlOrigin, UrlShort, 0).Err()
		if err != nil{
			return nil, err
		}

		return &Url{UrlOrigin: UrlOrigin, UrlShort: UrlShort}, nil
	} else if err != nil {
		return nil, err
	}
	return &Url{UrlOrigin: UrlOrigin, UrlShort: UrlShort}, nil
}


func (redis *Redis) GetByUrlShort(ctx context.Context, UrlShort string) (*Url, error){
	UrlOrigin, err := redis.client.Get(ctx, UrlShort).Result()
	if err == RedisLibrary.Nil{
		return nil, errors.New("not exist")
	}
	if err != nil {
		return nil, err
	}
	return &Url{UrlOrigin: UrlOrigin, UrlShort: UrlShort}, nil
}


func (redis *Redis) GetByUrlOrigin(ctx context.Context, UrlOrigin string) (*Url, error){
	UrlShort, err := redis.client.Get(ctx, UrlOrigin).Result()
	if err == RedisLibrary.Nil{
		return nil, errors.New("not exist")
	}
	if err != nil {
		return nil, err
	}
	return &Url{UrlOrigin: UrlOrigin, UrlShort: UrlShort}, nil
}
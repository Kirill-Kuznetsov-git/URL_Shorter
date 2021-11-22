package db

import (
	"URLShortener/app/config"
	"URLShortener/app/hasher"
	"context"
	"encoding/json"
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


func (redis *Redis) Save(ctx context.Context, UrlKey string, url ShortUrl) (string, error) {
	var id uint64
	counter, err := redis.client.Get(ctx, "counter").Result()
	if err == RedisLibrary.Nil{
		err := redis.client.Set(ctx, "counter", 1, 0).Err()
		if err != nil {
			return "", err
		}
		id = 1
	} else{
		redis.client.Incr(ctx, "counter")
		id, _ = strconv.ParseUint(counter, 10,64)
		id += 1
	}
	url.Id = id
	url_string := config.StructToString(url)

	res, err := redis.client.Get(ctx, UrlKey).Result()
	if err == RedisLibrary.Nil {
		log.Println("Redis: url will be created", url_string, res)

		err = redis.client.Set(ctx, UrlKey, url_string, 0).Err()
		if err != nil {
			return "Redis: set error. Please retry", err
		}
		return hasher.Encode(id), nil
	} else if err != nil {
		return "Redis: get error. Please retry", err
	}
	return "Redis: url already exists", errors.New("already exists")
}

func (redis *Redis) Get(ctx context.Context, UrlEncode string) (string, error){
	id, _ := hasher.Decode(UrlEncode)
	res, err := redis.client.Get(ctx, UrlEncode).Result()
	if err == RedisLibrary.Nil{
		return "Redis: Such url does not exists", errors.New("not exists")
	}
	TmpStruct := ShortUrl{}
	err = json.Unmarshal([]byte(res), &TmpStruct)
	if err != nil {
		return "", err
	}
	TmpStruct.Visits += 1

	redis.client.Set(ctx, UrlEncode, config.StructToString(TmpStruct), 0)
	return hasher.Encode(), nil
}
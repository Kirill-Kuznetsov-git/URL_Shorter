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


func (redis *Redis) Save(ctx context.Context, url ShortUrl) (string, error) {
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
	urlString := config.StructToString(url)

	res, err := redis.client.Get(ctx, url.UrlOrigin).Result()
	if err == RedisLibrary.Nil {
		log.Println("Redis: url will be created", urlString, res)

		err = redis.client.Set(ctx, url.UrlOrigin, urlString, 0).Err()
		if err != nil {
			return "Redis: set error. Please retry", err
		}
		err = redis.client.Set(ctx, strconv.FormatUint(id, 10), urlString, 0).Err()
		if err != nil {
			return "Redis: set error. Please retry", err
		}
		return hasher.Encode(id), nil
	} else if err != nil {
		return "Redis: get error. Please retry", err
	}
	TmpStruct := ShortUrl{}
	_ = json.Unmarshal([]byte(res), &TmpStruct)
	log.Println(hasher.Encode(TmpStruct.Id))
	return hasher.Encode(TmpStruct.Id), errors.New("already exist")
}


func (redis *Redis) Get(ctx context.Context, UrlShort string) (string, error){
	log.Println(UrlShort)
	id, _ := hasher.Decode(UrlShort)
	log.Println(id)
	res, err := redis.client.Get(ctx, strconv.FormatUint(id, 10)).Result()
	if err == RedisLibrary.Nil{
		log.Println(res)
		return "Redis: Such url does not exists", errors.New("not exists")
	}
	TmpStruct := ShortUrl{}
	err = json.Unmarshal([]byte(res), &TmpStruct)
	if err != nil {
		return "", err
	}
	TmpStruct.Visits += 1
	TmpStructString := config.StructToString(TmpStruct)
	redis.client.Set(ctx, strconv.FormatUint(id, 10), TmpStructString, 0)
	redis.client.Set(ctx, TmpStruct.UrlOrigin, TmpStructString, 0)
	return TmpStruct.UrlOrigin, nil
}
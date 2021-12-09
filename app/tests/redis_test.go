package tests

import (
	"URLShortener/app/db"
	"URLShortener/app/hasher"
	"context"
	"testing"
)

func TestInitRedis(t *testing.T){
	redis := db.Redis{}
	err := redis.Init()

	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	redis.Close()
}

func TestGetRedis(t *testing.T){
	redis := db.Redis{}
	err := redis.Init()
	ctx := context.Background()
	if err != nil {
		t.Errorf("Error with Init appeared")
	}
	defer redis.Close()
	hash, _ := hasher.Encode()
	UrlOrigin := "https::example.com"
	_, err = redis.Save(ctx, UrlOrigin)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	res, err := redis.GetByUrlShort(ctx, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Get appeared")
	}
	if res.UrlOrigin != UrlOrigin{
		t.Errorf("Different result from Get and Save")
	}
}

func TestCheckRedis(t *testing.T){
	redis := db.Redis{}
	err := redis.Init()
	ctx := context.Background()
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	defer redis.Close()
	UrlOrigin := "https::example.com"
	_, err = redis.Save(ctx, UrlOrigin)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	url, err := redis.GetByUrlOrigin(ctx, UrlOrigin)
	if url == nil && err.Error() == "not exist"{
		t.Errorf("Wrong answer from Check")
	}
}

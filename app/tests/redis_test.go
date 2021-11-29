package tests

import (
	"URLShortener/app/db"
	"URLShortener/app/hasher"
	"context"
	"testing"
)

func TestInitRedis(t *testing.T){
	redis, err := db.InitRedis()
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	if redis == nil{
		t.Errorf("Pointer to DB is nil")
	}
	redis.Close()
}

func TestGetRedis(t *testing.T){
	redis, err := db.InitRedis()
	ctx := context.Background()
	if err != nil {
		t.Errorf("Error with Init appeared")
	}
	defer redis.Close()
	if redis == nil{
		t.Errorf("Pointer to DB is nil")
	}
	hash, _ := hasher.Encode()
	UrlOrigin := "https::example.com"
	_, err = redis.Save(ctx, UrlOrigin, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	res, err := redis.Get(ctx, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Get appeared")
	}
	if res != UrlOrigin{
		t.Errorf("Different result from Get and Save")
	}
}

func TestCheckRedis(t *testing.T){
	redis, err := db.InitRedis()
	ctx := context.Background()
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	defer redis.Close()
	hash, _ := hasher.Encode()
	UrlOrigin := "https::example.com"
	_, err = redis.Save(ctx, UrlOrigin, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	flag, err := db.Check(ctx, UrlOrigin)
	if flag == "" && err.Error() == "not exist"{
		t.Errorf("Wrong answer from Check")
	}
}

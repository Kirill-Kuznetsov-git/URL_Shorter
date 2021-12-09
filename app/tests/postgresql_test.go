package tests

import (
	"URLShortener/app/db"
	"URLShortener/app/hasher"
	"context"
	"testing"
)

func TestInitPostgreSQL(t *testing.T){
	sql := db.PostgreSQL{}
	err := sql.Init()

	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	sql.Close()
}

func TestGetPostgreSQL(t *testing.T){
	sql := db.PostgreSQL{}
	err := sql.Init()
	ctx := context.Background()
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	defer sql.Close()
	hash, _ := hasher.Encode()
	UrlOrigin := "https::example.com"
	_, err = sql.Save(ctx, UrlOrigin)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	res, err := sql.GetByUrlShort(ctx, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Get appeared")
	}
	t.Log(res)
	t.Log(UrlOrigin)
	if res.UrlOrigin != UrlOrigin{
		t.Errorf("Different result from Get and Save")
	}
}

func TestCheckPostgreSQL(t *testing.T){
	sql := db.PostgreSQL{}
	err := sql.Init()
	ctx := context.Background()
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	defer sql.Close()
	UrlOrigin := "https::example.com"
	_, err = sql.Save(ctx, UrlOrigin)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	url, err := sql.GetByUrlOrigin(ctx, UrlOrigin)
	if url == nil && err.Error() == "not exist"{
		t.Errorf("Wrong answer from Check")
	}
}
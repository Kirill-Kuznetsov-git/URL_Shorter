package tests

import (
	"URLShortener/app/db"
	"URLShortener/app/hasher"
	"context"
	"testing"
)

func TestInitPostgreSQL(t *testing.T){
	sql, err := db.InitPostgreSQL()

	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	if sql == nil{
		t.Errorf("Pointer to DB is nil")
	}
	sql.Close()
}

func TestGetPostgreSQL(t *testing.T){
	sql, err := db.InitPostgreSQL()
	ctx := context.Background()
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	defer sql.Close()
	if sql == nil{
		t.Errorf("Pointer to DB is nil")
	}
	hash, _ := hasher.Encode()
	UrlOrigin := "https::example.com"
	_, err = sql.Save(ctx, UrlOrigin, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	res, err := sql.Get(ctx, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Get appeared")
	}
	t.Log(res)
	t.Log(UrlOrigin)
	if res != UrlOrigin{
		t.Errorf("Different result from Get and Save")
	}
}

func TestCheckPostgreSQL(t *testing.T){
	sql, err := db.InitPostgreSQL()
	ctx := context.Background()
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Init appeared")
	}
	defer sql.Close()
	hash, _ := hasher.Encode()
	UrlOrigin := "https::example.com"
	_, err = sql.Save(ctx, UrlOrigin, hash)
	if err != nil {
		t.Log(err.Error())
		t.Errorf("Error with Save appeared")
	}
	flag, err := db.Check(ctx, UrlOrigin)
	if flag == "" && err.Error() == "not exist"{
		t.Errorf("Wrong answer from Check")
	}
}
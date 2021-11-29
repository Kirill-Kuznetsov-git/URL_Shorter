package db

import (
	"URLShortener/app/hasher"
	"context"
	"errors"
	"log"
)

type DB struct {
	postgre *PostgreSQL
	redis *Redis
}

var dbUniversal DB

func InitDB(nameDb string) (*DB, error){
	log.Println("Database name from configurations: " + nameDb)
	switch nameDb {
	case "redis":
		redis, err := InitRedis()
		if err != nil {
			log.Println("DB ERROR")
			return nil, err
		}
		dbUniversal.redis = redis
	case "postgreSQL":
		sql, err := InitPostgreSQL()
		if err != nil {
			log.Println("DB ERROR")
			return nil, err
		}
		dbUniversal.postgre = sql
	}
	return &dbUniversal, nil
}

func (db *DB) Close() error{
	if db.postgre != nil{
		db.postgre.Close()
	}
	if db.redis != nil{
		err := db.redis.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func Save(ctx context.Context, UrlOrigin string) (string, error){
	log.Println("Come to Save function in db_universal")

	url, err := Check(ctx, UrlOrigin)
	if err == nil{
		return url, nil
	} else if err.Error() != "not exist"{
		return "error", err
	}
	UrlShort, _ := hasher.Encode()
	_, err = Get(ctx, UrlShort)
	for err == nil || (err != nil && err.Error() != "not exist"){
		log.Println(err)
		UrlShort, _ = hasher.Encode()
		_, err = Get(ctx, UrlShort)
	}

	if dbUniversal.postgre != nil {
		url, err := dbUniversal.postgre.Save(ctx, UrlOrigin, UrlShort)
		if err != nil {
			return "postgre error in Save", err
		}
		return url, nil
	} else if dbUniversal.redis != nil{
		url, err := dbUniversal.redis.Save(ctx, UrlOrigin, UrlShort)
		if err != nil{
			if err.Error() == "already exist"{
				return url, err
			}
			return "error", err
		}

		return url, err
	}
	return "Wrong DB", errors.New("wrong db")
}

func Get(ctx context.Context, UrlShort string) (string, error){
	if dbUniversal.postgre != nil {
		url, err := dbUniversal.postgre.Get(ctx, UrlShort)
		if err != nil {
			return "postgre error", err
		}
		return url, nil
	} else if dbUniversal.redis != nil{
		url, err := dbUniversal.redis.Get(ctx, UrlShort)
		if err != nil {
			return "error", err
		}
		return url, nil
	}
	return "Wrong DB", errors.New("wrong db")
}

func Check(ctx context.Context, UrlOrigin string) (string, error){
	if dbUniversal.postgre != nil {
		url, err := dbUniversal.postgre.Check(ctx, UrlOrigin)
		return url, err
	} else if dbUniversal.redis != nil{
		url, err := dbUniversal.redis.Check(ctx, UrlOrigin)
		return url, err
	}
	return "error", errors.New("wrong db")
}

package db

import (
	"context"
	"errors"
	"log"
)

type DB struct {
	postgre *PostgreSQL
	redis *Redis
}

var dbUniversal DB

func InitDB(nameDb string) *DB{
	log.Println("Database name from .env: " + nameDb)
	switch nameDb {
	case "redis":
		dbUniversal.redis, _ = InitRedis()
	case "postgreSQL":
		log.Println("POSTGRE IS NOT EMPTY")
		dbUniversal.postgre, _ = InitPostgreSQL()
	}
	return &dbUniversal
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

func Save(ctx context.Context, url string) (string, error){
	log.Println("Come to Save function in db_universal")
	if dbUniversal.postgre != nil {
		log.Println("Come to postgreSQL")
		url, err := dbUniversal.postgre.Save(ctx, url)
		if err != nil {
			return "postgre error", err
		}
		return url, nil
	} else if dbUniversal.redis != nil{
		url, err := dbUniversal.redis.Save(ctx, url)
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
		url, err := dbUniversal.postgre.Save(ctx, UrlShort)
		if err != nil {
			return "postgre error", err
		}
		return url, nil
	} else if dbUniversal.redis != nil{
		url, err := dbUniversal.redis.Get(ctx, UrlShort)
		if err != nil {
			return "error", err
		}
		return url, err
	}
	return "Wrong DB", errors.New("wrong db")
}
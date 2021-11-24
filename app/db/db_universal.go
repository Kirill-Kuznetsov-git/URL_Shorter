package db

import (
	"context"
	"errors"
)

type DB struct {
	postgre *PostgreSQL
	redis *Redis
}

var dbUniversal DB

func InitDB(nameDb string) *DB{
	switch nameDb {
	case "redis":
		dbUniversal.redis, _ = InitRedis()
	case "postgreSQL":
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

func Save(ctx context.Context, url ShortUrl) (string, error){
	if dbUniversal.postgre != nil {
		dbUniversal.postgre.Save()
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
		dbUniversal.postgre.Save()
	} else if dbUniversal.redis != nil{
		url, err := dbUniversal.redis.Get(ctx, UrlShort)
		if err != nil {
			return "error", err
		}
		return url, err
	}
	return "Wrong DB", errors.New("wrong db")
}
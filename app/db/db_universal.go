package db

import (
	"context"
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
	case "postgre":
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

func GetDB() *DB{
	return &dbUniversal
}

func Save(ctx context.Context, url ShortUrl) (string, error){
	if dbUniversal.postgre != nil{
		dbUniversal.postgre.Save()
	} else if dbUniversal.redis != nil{
		url, err := dbUniversal.redis.Save(ctx, url)
		if err != nil {
			return "error", err
		}
		return url, err
	}
	return "qwe", nil
}
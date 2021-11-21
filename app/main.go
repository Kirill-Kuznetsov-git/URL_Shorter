package main

import (
	"URLShortener/app/db"
	"log"
)

func main() {
	postgre, err := db.InitPostgreSQL()
	if err != nil {
		log.Fatalf("Could not initialize Database connection %s", err)
	}
	defer postgre.Close()

	redis, err := db.InitRedis()
	if err != nil {
		log.Fatalf("Could not initialize Redis client %s", err)
	}
	log.Println(redis)
}

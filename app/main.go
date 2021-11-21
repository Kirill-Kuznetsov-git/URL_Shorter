package main

import (
	"URLShortener/app/db"
	"fmt"
	"log"
)

func main() {
	fmt.Println("ASD")
	postgre, err := db.InitPostgreSQL()
	if err != nil {
		log.Fatalf("Could not initialize Database connection %s", err)
	}
	defer postgre.Close()
	fmt.Println("ASD")
	redis, err := db.InitRedis()
	if err != nil {
		log.Fatalf("Could not initialize Redis client %s", err)
	}
	log.Println(redis)
}

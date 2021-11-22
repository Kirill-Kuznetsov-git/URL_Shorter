package main

import (
	"github.com/gorilla/mux"
	"URLShortener/app/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	port := "8080"
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

	log.Fatal(http.ListenAndServe(":" + port, router))
}

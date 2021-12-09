package main

import (
	"URLShortener/app/config"
	"URLShortener/app/controllers"
	DBpackage "URLShortener/app/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	configuration, err := config.FromFile("./configurations.json")
	if err != nil {
		log.Fatal(err)
	}
	if configuration.Db.DbName == "redis"{
		DBpackage.Db = &DBpackage.Redis{}
	} else if configuration.Db.DbName == "postgreSQL"{
		DBpackage.Db = &DBpackage.PostgreSQL{}
	}
	err = DBpackage.Db.Init()
	if err != nil{
		log.Fatal(err)
	}
	defer DBpackage.Db.Close()

	router.HandleFunc("/create_url", controllers.CreateURL).Methods("POST")
	router.HandleFunc("/{shortLink}", controllers.Redirect).Methods("GET")

	log.Fatal(http.ListenAndServe(":" + configuration.Server.Port, router))
}

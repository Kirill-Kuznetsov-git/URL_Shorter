package config

import (
	"URLShortener/app/config"
	DBpackage "URLShortener/app/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	configuration, err := config.FromFile("./configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	db := DBpackage.InitDB(configuration.Db.DbName)
	log.Println(db)

	log.Fatal(http.ListenAndServe(":" + configuration.Server.Port, router))
}

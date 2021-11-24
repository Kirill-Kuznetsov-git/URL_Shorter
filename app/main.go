package main

import (
	"URLShortener/app/controllers"
	DBpackage "URLShortener/app/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	db := DBpackage.InitDB(os.Getenv("DATABASE_NAME"))
	defer db.Close()

	router.HandleFunc("/create_url", controllers.CreateURL).Methods("POST")
	router.HandleFunc("/{shortLink}", controllers.Redirect).Methods("GET")

	log.Fatal(http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), router))
}

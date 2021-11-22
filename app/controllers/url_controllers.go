package controllers

import (
	"URLShortener/app/config"
	dbpackage "URLShortener/app/db"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateURL = func(w http.ResponseWriter, r *http.Request){
	fmt.Println("HELLOO")
	URLstruct := &dbpackage.ShortUrl{}
	err := json.NewDecoder(r.Body).Decode(URLstruct)
	if err != nil {
		fmt.Println("Error with json")
		return
	}
	save, err := dbpackage.Save(r.Context(), *URLstruct)
	if err != nil {
		return
	}
	config.Respond(w, save)

}

var GetURL = func(w http.ResponseWriter, r *http.Request){
	fmt.Println("Bue")
}
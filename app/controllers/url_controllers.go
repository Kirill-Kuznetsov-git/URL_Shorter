package controllers

import (
	"URLShortener/app/config"
	dbpackage "URLShortener/app/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var CreateURL = func(w http.ResponseWriter, r *http.Request){
	URLstruct := &dbpackage.Url{}
	err := json.NewDecoder(r.Body).Decode(URLstruct)
	if err != nil {
		fmt.Println("Error with json")
		return
	}
	save, err := dbpackage.Save(r.Context(), (*URLstruct).UrlOrigin)
	if err != nil {
		if err.Error() == "already exist"{
			config.Respond(w, save)
		}
		log.Println(err)
		return
	}
	config.Respond(w, save)

}

var Redirect = func(w http.ResponseWriter, r *http.Request){
	UrlShort := r.RequestURI[1:]
	UrlOrigin, err := dbpackage.Get(r.Context(), UrlShort)
	log.Println("HIHIHI")
	log.Println("UrlOrigin1: " + UrlOrigin)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println("UrlOrigin2: " + UrlOrigin)
	http.Redirect(w, r, UrlOrigin, 301)
}

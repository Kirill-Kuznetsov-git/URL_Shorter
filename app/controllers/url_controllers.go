package controllers

import (
	"URLShortener/app/config"
	dbpackage "URLShortener/app/db"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateURL = func(w http.ResponseWriter, r *http.Request){
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

var Redirect = func(w http.ResponseWriter, r *http.Request){
	UrlShort := r.RequestURI[1:]
	UrlOrigin, err := dbpackage.Get(r.Context(), UrlShort)
	if err != nil {
		return
	}
	http.Redirect(w, r, UrlOrigin, 301)
}

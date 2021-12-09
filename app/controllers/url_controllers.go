package controllers

import (
	"URLShortener/app/config"
	dbpackage "URLShortener/app/db"
	"encoding/json"
	"log"
	"net/http"
)

// CreateURL Handler for creating new URL
// In the body must be {"url_origin": <string>}.
// Will be created Short Url for this Url origin, using hasher
var CreateURL = func(w http.ResponseWriter, r *http.Request){
	URLstruct := &dbpackage.Url{}
	err := json.NewDecoder(r.Body).Decode(URLstruct)
	if err != nil {
		log.Println("Error with json")
		return
	}
	save, err := dbpackage.Db.Save(r.Context(), (*URLstruct).UrlOrigin)
	if err != nil {
		if err.Error() == "already exist"{
			config.Respond(w, save)
		}
		log.Println(err.Error())
		return
	}
	config.Respond(w, save)

}
// Redirect Handler which get Short Url from search bar,
// Find in the DB Origin Url and redirect to this Origin Url
var Redirect = func(w http.ResponseWriter, r *http.Request){
	UrlShort := r.RequestURI[1:]
	res, err := dbpackage.Db.GetByUrlShort(r.Context(), UrlShort)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	http.Redirect(w, r, res.UrlOrigin, 301)
}

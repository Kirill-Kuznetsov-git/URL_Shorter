package controllers

import (
	"URLShortener/app/config"
	dbpackage "URLShortener/app/db"
	"encoding/json"
	"log"
	"net/http"
)

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

var Redirect = func(w http.ResponseWriter, r *http.Request){
	UrlShort := r.RequestURI[1:]
	res, err := dbpackage.Db.GetByUrlShort(r.Context(), UrlShort)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	http.Redirect(w, r, res.UrlOrigin, 301)
}

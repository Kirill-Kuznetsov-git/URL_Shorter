package controllers

import (
	"fmt"
	"net/http"
)

var CreateURL = func(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hi")
}

var GetURL = func(w http.ResponseWriter, r *http.Request){
	fmt.Println("Bue")
}
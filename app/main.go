package main

import (

	"URLShortener/app/hasher"
	"fmt"

)

func main() {
	fmt.Println(hasher.Encode(10000000000000000000))
	//router := mux.NewRouter()
	//
	//configuration, err := config.FromFile("./configuration.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//db := DBpackage.InitDB(configuration.Db.DbName)
	//defer db.Close()
	//
	//router.HandleFunc("/get_url", controllers.GetURL).Methods("GET")
	//router.HandleFunc("/create_url", controllers.CreateURL).Methods("POST")
	//
	//log.Fatal(http.ListenAndServe(":" + configuration.Server.Port, router))
}

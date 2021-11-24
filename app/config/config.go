package config

import (
	"encoding/json"
	"log"
	"net/http"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Db struct {
		DbName string `json:"db_name"`
	} `json:"db"`
}


func StructToString(SomeStruct interface{}) string {
	res, err := json.Marshal(SomeStruct)
	if err != nil {
		log.Println(err)
	}

	return string(res)
}

func Respond(w http.ResponseWriter, data interface{})  {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return 
	}
}
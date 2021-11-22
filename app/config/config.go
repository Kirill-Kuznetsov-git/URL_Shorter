package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Db struct {
		DbName string `json:"db_name"`
	} `json:"db"`
}

func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}


func StructToString(SomeStruct interface{}) string {
	res, err := json.Marshal(SomeStruct)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(res))

	return string(res)
}
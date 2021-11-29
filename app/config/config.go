package config

import (
	"encoding/json"
	"io/ioutil"
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

func Respond(w http.ResponseWriter, data interface{})  {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return 
	}
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
package db

import "time"

type Service interface {
	Save(string, time.Time) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*Item, error)
	Close() error
}

type Item struct {
	Id      uint64 `json:"id" redis:"id" gorm:"id"`
	URL     string `json:"url" redis:"url" gorm:"url"`
	Expires string `json:"expires" redis:"expires" gorm:"expires"`
	Visits  int    `json:"visits" redis:"visits" gorm:"visits"`
}

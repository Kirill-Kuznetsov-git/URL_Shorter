package db

import "time"

type Service interface {
	Save(string, time.Time) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*ShortUrl, error)
	Close() error
}

type ShortUrl struct {
	Id uint64 `json:"id" redis:"id" pgxpool:"id"`
	Visits     int    `json:"visits" redis:"visits" pgxpool:"visits"`
}

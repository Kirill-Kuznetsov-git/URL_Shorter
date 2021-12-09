package db

import "context"

var Db dbInterface

type dbInterface interface {
	Init() error
	Close() error
	Save(ctx context.Context, UrlOrigin string) (*Url, error)
	GetByUrlShort(ctx context.Context, UrlShort string) (*Url, error)
	GetByUrlOrigin(ctx context.Context, UrlOrigin string) (*Url, error)
}


type Url struct{
	UrlShort string `json:"url_short" redis:"url_short" pgxpool:"url_short"`
	UrlOrigin string `json:"url_origin" redis:"url_origin" pgxpool:"url_origin"`
}

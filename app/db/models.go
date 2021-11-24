package db

type ShortUrl struct {
	Id uint64 `json:"id" redis:"id" pgxpool:"id"`
	UrlOrigin string `json:"url_origin" redis:"url_origin" pgxpool:"id_origin"`
	Visits     int    `json:"visits" redis:"visits" pgxpool:"visits"`
}

package db

type Url struct{
	UrlShort string `json:"url_short" redis:"url_short" pgxpool:"url_short"`
	UrlOrigin string `json:"url_origin" redis:"url_origin" pgxpool:"url_origin"`
}

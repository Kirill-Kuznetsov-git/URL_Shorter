package db

import (
	"URLShortener/app/hasher"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

type PostgreSQL struct {
	pool *sql.DB
}

func (postgre *PostgreSQL) Init() error {
	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	if err != nil {
		dbPort = 5432
	}
	pool, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	log.Println("Successfully connected to postgreSQL")
	log.Println(pool)
	postgre.pool = pool
	return nil
}

func (postgre *PostgreSQL) Close() error {
	postgre.pool.Close()
	return nil
}


func (postgre *PostgreSQL) Save(ctx context.Context, UrlOrigin string) (*Url, error){
	url, err := postgre.GetByUrlOrigin(ctx, UrlOrigin)
	if err == nil{
		return url, nil
	} else if err.Error() != "not exist"{
		return nil, err
	}

	UrlShort, _ := hasher.Encode()
	_, err = postgre.GetByUrlShort(ctx, UrlShort)

	for err == nil || err.Error() != "not exist"{
		UrlShort, _ = hasher.Encode()
		_, err = postgre.GetByUrlShort(ctx, UrlShort)
	}
	query := `INSERT INTO url (url_short, url_origin) VALUES($1, $2)`
	_ = postgre.pool.QueryRowContext(ctx, query, UrlShort, UrlOrigin)
	log.Println("Result UrlShort from postgreSQL: " + UrlShort)
	return &Url{UrlShort: UrlShort, UrlOrigin: UrlOrigin}, nil
}

func (postgre *PostgreSQL) GetByUrlShort(ctx context.Context, UrlShort string) (*Url, error){
	query := "SELECT url_origin FROM url WHERE url_short=$1"

	res := Url{UrlShort: UrlShort}
	row := postgre.pool.QueryRowContext(ctx, query, UrlShort)
	err := row.Scan(&res.UrlOrigin)
	if err != nil{
		if err.Error() == "sql: no rows in result set"{
			return nil, errors.New("not exist")
		}
		return nil, err
	}
	return &res, nil
}

func (postgre *PostgreSQL) GetByUrlOrigin(ctx context.Context, UrlOrigin string) (*Url, error){
	query := "SELECT url_short FROM url WHERE url_origin=$1"

	res := Url{UrlOrigin: UrlOrigin}
	row := postgre.pool.QueryRowContext(ctx, query, UrlOrigin)
	err := row.Scan(&res.UrlShort)
	if err != nil{
		if err.Error() == "sql: no rows in result set"{
			return nil, errors.New("not exist")
		}
		return nil, err
	}

	return &res, nil
}
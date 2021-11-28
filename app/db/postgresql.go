package db

import (
	"URLShortener/app/hasher"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

type PostgreSQL struct {
	pool *sql.DB
}

func InitPostgreSQL() (*PostgreSQL, error) {
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
		fmt.Println(err.Error())
		panic(err)
	}

	err = pool.Ping()
	for err != nil{
		err = pool.Ping()
	}
	fmt.Println("Successfully connected to postgreSQL")
	return &PostgreSQL{
		pool: pool,
	}, nil
}

func (p *PostgreSQL) Close() {
	p.pool.Close()
}


func (p *PostgreSQL) Save(ctx context.Context, UrlOrigin string) (string, error){
	query := `INSERT INTO url (url_short, url_origin) VALUES($1, $2)`
	log.Println("I AM HERE")
	UrlShort, err := hasher.Encode()
	if err != nil{
		return "hasher error", err
	}

	log.Println("I AM HERE 1")
	_ = p.pool.QueryRowContext(ctx, query, UrlShort, UrlOrigin)
	log.Println("I AM HERE 2")
	log.Println("Result UrlShort from postgreSQL: " + UrlShort)
	return UrlShort, nil
}

func (p *PostgreSQL) Get(ctx context.Context, UrlShort string) (string, error){
	query := "SELECT url_origin FROM url WHERE url_short=?"

	var res Url
	row := p.pool.QueryRowContext(ctx, query, UrlShort)
	err := row.Scan(&res.UrlOrigin)
	if err != nil{
		return "postgre error", err
	}

	return res.UrlOrigin, nil
}
package db

import (
	"URLShortener/app/hasher"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"strconv"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func InitPostgreSQL() (*PostgreSQL, error) {
	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		dbPort = 5432
	}
	dbUri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=verify-ca&pool_max_conns=10", dbUsername, dbPassword, dbHost, dbPort, dbName)
	pool, err := pgxpool.Connect(context.Background(), dbUri)
	log.Println("POSTGREURL: " + dbUri)
	if err != nil {
		log.Println("PANIC PANIC WRONG DB")
		return nil, err
	}

	return &PostgreSQL{
		pool: pool,
	}, nil
}

func (p *PostgreSQL) Close() {
	p.pool.Close()
}


func (p *PostgreSQL) Save(ctx context.Context, UrlOrigin string) (string, error){
	query := `INSERT INTO URL (url_short, url_origin) VALUES($1, $2)`

	var res Url
	UrlShort, err := hasher.Encode()
	if err != nil{
		return "hasher error", err
	}
	if err := p.pool.QueryRow(ctx, query, UrlShort, UrlOrigin).
		Scan(&res.UrlOrigin); err != nil {
		return "postgre error", err
	}
	log.Println("Result UrlShort from postgreSQL: " + res.UrlShort)
	return res.UrlShort, nil
}

func (p *PostgreSQL) Get(ctx context.Context, UrlShort string) (string, error){
	query := `SELECT url_origin FROM "URL" WHERE url_short = $1`

	var res Url

	if err := p.pool.QueryRow(ctx, query, UrlShort).
		Scan(&res.UrlOrigin); err != nil {
		return "postgre error", err
	}

	return res.UrlOrigin, nil
}
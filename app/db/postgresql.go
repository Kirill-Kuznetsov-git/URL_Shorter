package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strconv"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func InitPostgreSQL() (*PostgreSQL, error) {
	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		dbPort = 5432
	}
	dbUri := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUsername, dbName, dbPassword)
	pool, err := pgxpool.Connect(context.Background(), dbUri)
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{
		pool: pool,
	}, nil
}

func (p *PostgreSQL) Close() {
	p.pool.Close()
}

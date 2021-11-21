package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func InitPostgreSQL() (*PostgreSQL, error) {
	fmt.Println("fd")
	fmt.Println(os.Getenv("DATABASE_URL"))
	fmt.Println("q")
	pool, err := pgxpool.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/app")
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

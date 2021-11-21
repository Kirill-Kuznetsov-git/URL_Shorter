package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func InitPostgreSQL() (*PostgreSQL, error) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
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

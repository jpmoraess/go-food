package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	connString = "postgres://backend:backend@localhost:5432/postgres?sslmode=disable"
)

func CreateConnectionPool() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("unable to parse DATABASE_URL: %v\n", err)
	}

	// additional configs
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Second

	// pool
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v\n", err)
		return nil, err
	}
	return dbpool, nil
}

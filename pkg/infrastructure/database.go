package infrastructure

import (
	"context"
	"github.com/jackc/pgx/pgxpool"
)

const dsn = "host=localhost port=5432 user=postgres password=postgres dbname=k_on pool_max_conns=100"

func InitDatabaseConnection() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	return pgxpool.ConnectConfig(context.Background(), config)
}

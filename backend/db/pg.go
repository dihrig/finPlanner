package db

import (
    "context"
    "fmt"

    "github.com/dihrig/finPlanner/config"
    "github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg config.Config) (*pgxpool.Pool, error) {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
    )
    return pgxpool.New(context.Background(), dsn)
}

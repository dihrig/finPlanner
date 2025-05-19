package db

import (
	"context"
	"fmt"

	"github.com/dihrig/finPlanner/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name,
	)
	return pgxpool.New(context.Background(), dsn)
}

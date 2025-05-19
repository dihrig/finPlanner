package handlers

import (
    "context"
    "net/http"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
)

func HealthCheck(db *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
        defer cancel()

        if err := db.Ping(ctx); err != nil {
            http.Error(w, "Database unreachable", http.StatusInternalServerError)
            return
        }

        w.Write([]byte("Database connection OK"))
    }
}

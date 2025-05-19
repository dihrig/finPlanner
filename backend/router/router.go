package router

import (
    "net/http"

    "github.com/dihrig/finPlanner/handlers"
    "github.com/go-chi/chi/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

func New(db *pgxpool.Pool) http.Handler {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })

    r.Get("/db", handlers.HealthCheck(db))

    return r
}

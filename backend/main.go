package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/dihrig/finPlanner/config"
    "github.com/dihrig/finPlanner/db"
    "github.com/dihrig/finPlanner/router"
)

func main() {
    cfg := config.Load()

    pool, err := db.Connect(cfg)
    if err != nil {
        log.Fatalf("DB error: %v", err)
    }
    defer pool.Close()

    r := router.New(pool)

    fmt.Println("Server started on :" + cfg.Port)
    http.ListenAndServe(":"+cfg.Port, r)
}

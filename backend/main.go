package main

import (
	"log"
	"net/http"

	"github.com/dihrig/finPlanner/config"
	"github.com/dihrig/finPlanner/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	conn, err := db.Connect(&cfg.DB)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer conn.Close()

	log.Printf("Server starting on :%s...", cfg.Server.Port)
	http.ListenAndServe(":"+cfg.Server.Port, nil)
}

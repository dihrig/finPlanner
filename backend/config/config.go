// config/config.go
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type ServerConfig struct {
	Port string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	requiredEnv := []string{
		"POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB",
		"POSTGRES_HOST", "POSTGRES_PORT", "BACKEND_PORT",
	}

	for _, key := range requiredEnv {
		if os.Getenv(key) == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", key)
		}
	}

	return &Config{
		DB: DBConfig{
			User: os.Getenv("POSTGRES_USER"),
			Pass: os.Getenv("POSTGRES_PASSWORD"),
			Host: os.Getenv("POSTGRES_HOST"),
			Port: os.Getenv("POSTGRES_PORT"),
			Name: os.Getenv("POSTGRES_DB"),
		},
		Server: ServerConfig{
			Port: os.Getenv("BACKEND_PORT"),
		},
	}, nil
}

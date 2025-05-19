package config

import "os"

type Config struct {
    Port     string
    DBUser   string
    DBPass   string
    DBHost   string
    DBPort   string
    DBName   string
}

func Load() Config {
    return Config{
        Port:     getEnv("PORT", "8080"),
        DBUser:   getEnv("DB_USER", "myuser"),
        DBPass:   getEnv("DB_PASSWORD", "mypassword"),
        DBHost:   getEnv("DB_HOST", "db"),
        DBPort:   getEnv("DB_PORT", "5432"),
        DBName:   getEnv("DB_NAME", "mydatabase"),
    }
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}

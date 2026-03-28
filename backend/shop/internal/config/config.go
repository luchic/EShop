package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	HTTPAddr         string
	TokenSecret      string
	ConnectionString string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	IdleTimeout      time.Duration
}

// TODO: AUTH_DATABASE_URL
func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:         getEnv("AUTH_HTTP_ADDR", ":8081"),
		TokenSecret:      getEnv("AUTH_TOKEN_SECRET", "dev-secret-change-me"),
		ConnectionString: getEnv("AUTH_DATABASE_URL", ""),
		ReadTimeout:      5 * time.Second,
		WriteTimeout:     5 * time.Second,
		IdleTimeout:      30 * time.Second,
	}

	if cfg.ConnectionString == "" {
		return Config{}, fmt.Errorf("AUTH_DATABASE_URL is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

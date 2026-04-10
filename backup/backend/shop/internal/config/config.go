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
	FortyTwoClientID string
	FortyTwoSecret   string
	FortyTwoRedirect string
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
		FortyTwoClientID: getEnv("AUTH_42_CLIENT_ID", ""),
		FortyTwoSecret:   getEnv("AUTH_42_CLIENT_SECRET", ""),
		FortyTwoRedirect: getEnv("AUTH_42_REDIRECT_URI", ""),
		ReadTimeout:      5 * time.Second,
		WriteTimeout:     5 * time.Second,
		IdleTimeout:      30 * time.Second,
	}

	if cfg.ConnectionString == "" {
		return Config{}, fmt.Errorf("AUTH_DATABASE_URL is required")
	}
	if cfg.FortyTwoClientID == "" {
		return Config{}, fmt.Errorf("AUTH_42_CLIENT_ID is required")
	}
	if cfg.FortyTwoSecret == "" {
		return Config{}, fmt.Errorf("AUTH_42_CLIENT_SECRET is required")
	}
	if cfg.FortyTwoRedirect == "" {
		return Config{}, fmt.Errorf("AUTH_42_REDIRECT_URI is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

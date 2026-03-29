package config

import "os"

type Config struct {
	Address     string
	DatabaseURL string
	FrontendURL string
}

func FromEnv() Config {
	return Config{
		Address:     envOrDefault("API_ADDRESS", ":8080"),
		DatabaseURL: envOrDefault("DATABASE_URL", "postgres://nexus:nexus@localhost:5432/nexus?sslmode=disable"),
		FrontendURL: envOrDefault("FRONTEND_URL", "http://localhost:3000"),
	}
}

func envOrDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port            string
	Environment     string
	DatabaseURL     string
	FirebaseProject string
	AllowedOrigins  string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

func Load() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		Environment:     getEnv("ENVIRONMENT", "development"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		FirebaseProject: getEnv("FIREBASE_PROJECT_ID", ""),
		AllowedOrigins:  getEnv("ALLOWED_ORIGINS", "*"),
		ReadTimeout:     getDuration("READ_TIMEOUT", 10),
		WriteTimeout:    getDuration("WRITE_TIMEOUT", 30),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getDuration(key string, fallback int) time.Duration {
	if val := os.Getenv(key); val != "" {
		if sec, err := strconv.Atoi(val); err == nil {
			return time.Duration(sec) * time.Second
		}
	}
	return time.Duration(fallback) * time.Second
}

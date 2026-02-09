// Package config handles loading configuration from environment variables.
// It uses the godotenv package to load values from a .env file if present,
// falling back to default values if not set.
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration values for the application.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// Load reads configuration from environment variables.
// It attempts to load from .env file first, then uses environment variables or defaults.
func Load() *Config {
	// Load .env file if it exists (ignore error if not found)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables or defaults")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "postsdb"),
	}
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

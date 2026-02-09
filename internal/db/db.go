// Package db handles database connection and initialization.
// It provides functions to establish a connection to PostgreSQL and set up the schema.
package db

import (
	"database/sql"
	"fmt"
	"log"

	"go-docker-app/internal/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global database connection pool.
// It's initialized once and used throughout the application.
var DB *sql.DB

// InitDB establishes a connection to the PostgreSQL database using the provided configuration.
// It also creates the necessary tables if they don't exist.
func InitDB(cfg *config.Config) {
	// Build connection string from config
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	// Open database connection
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Successfully connected to database")

	// Create tables
	createTable()
}

// createTable creates the posts table if it doesn't already exist.
// This ensures the database schema is set up on application startup.
func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create posts table:", err)
	}

	log.Println("Posts table created or already exists")
}

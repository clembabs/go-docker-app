// Package main is the entry point for the Go REST API server.
// It initializes the application, sets up routing, and starts the HTTP server.
package main

import (
	"log"
	"net/http"

	"go-docker-app/internal/config"
	"go-docker-app/internal/db"
	"go-docker-app/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from environment variables
	cfg := config.Load()

	// Initialize database connection and create tables
	db.InitDB(cfg)

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")

	// Log server startup
	log.Println("Server starting on :8080")

	// Start HTTP server on port 8080
	// This will block until the server is stopped
	log.Fatal(http.ListenAndServe(":8080", r))
}

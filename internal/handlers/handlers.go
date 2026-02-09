// Package handlers contains HTTP request handlers for the REST API.
// These functions handle incoming requests, interact with the database through the db package,
// and return appropriate responses.
package handlers

import (
	"encoding/json"
	"net/http"

	"go-docker-app/internal/db"
	"go-docker-app/internal/models"
)

// CreatePost handles POST /posts requests.
// It expects a JSON body with title and body fields, creates a new post in the database,
// and returns the created post with its ID and creation timestamp.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	// Decode JSON request body into Post struct
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Insert post into database and get back the ID and created_at
	query := "INSERT INTO posts (title, body) VALUES ($1, $2) RETURNING id, created_at"
	err := db.DB.QueryRow(query, post.Title, post.Body).Scan(&post.ID, &post.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to create post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers and encode post as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(post)
}

// GetPosts handles GET /posts requests.
// It retrieves all posts from the database and returns them as a JSON array.
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Query all posts from database
	rows, err := db.DB.Query("SELECT id, title, body, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Failed to retrieve posts: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []models.Post

	// Iterate through result rows and build posts slice
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt); err != nil {
			http.Error(w, "Failed to scan post: "+err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	// Check for any error that occurred during iteration
	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating posts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers and encode posts array as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

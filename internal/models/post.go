// Package models contains data structures used throughout the application.
// These represent the core business entities.
package models

import "time"

// Post represents a blog post or similar content item.
// It includes an ID, title, body content, and creation timestamp.
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

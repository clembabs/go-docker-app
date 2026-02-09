# Go REST API

A production-style Go REST API for managing posts, designed to be containerized with Docker and deployed to AWS EC2 behind Nginx, connected to RDS PostgreSQL.

## Features

- RESTful API with POST and GET endpoints for posts
- PostgreSQL database integration
- Environment variable configuration
- Automatic table creation on startup
- Clean project structure following Go best practices

## Project Structure

- `cmd/server/`: Main application entry point
- `internal/config/`: Configuration loading from environment variables
- `internal/db/`: Database connection and initialization
- `internal/handlers/`: HTTP request handlers
- `internal/models/`: Data models (Post struct)

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose

## Setup

1. **Clone or navigate to the project directory**

2. **Start PostgreSQL with Docker Compose:**

   ```bash
   docker-compose up -d
   ```

   This will start a PostgreSQL container on port 5432 with the database `postsdb`.

3. **Configure environment variables:**

   Copy the example environment file:

   ```bash
   cp .env.example .env
   ```

   The default values in `.env.example` should work with the Docker Compose setup. Adjust if needed.

4. **Install dependencies:**

   ```bash
   go mod tidy
   ```

5. **Run the application:**

   ```bash
   go run cmd/server/main.go
   ```

   The server will start on port 8080.

## API Endpoints

### POST /posts

Create a new post.

**Request Body:**
```json
{
  "title": "Post Title",
  "body": "Post content here"
}
```

**Response:**
```json
{
  "id": 1,
  "title": "Post Title",
  "body": "Post content here",
  "created_at": "2023-01-01T12:00:00Z"
}
```

### GET /posts

Retrieve all posts.

**Response:**
```json
[
  {
    "id": 1,
    "title": "Post Title",
    "body": "Post content here",
    "created_at": "2023-01-01T12:00:00Z"
  }
]
```

## Testing the API

You can test the endpoints using curl or tools like Postman:

```bash
# Create a post
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "My First Post", "body": "This is the content of my first post."}'

# Get all posts
curl http://localhost:8080/posts
```

## Docker Deployment (Future)

This project is structured to be easily containerized. For production deployment:

1. Create a Dockerfile
2. Use docker-compose for the full stack (app + database)
3. Deploy to AWS EC2 with Nginx as reverse proxy
4. Use RDS PostgreSQL instead of local container

## Environment Variables

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: password)
- `DB_NAME`: Database name (default: postsdb)
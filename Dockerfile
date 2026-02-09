# ---------- Stage 1: Build the Go binary ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (needed for go modules sometimes)
RUN apk add --no-cache git

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server


# ---------- Stage 2: Minimal runtime image ----------
FROM alpine:latest

WORKDIR /app

# Add certificates (important for HTTPS calls later)
RUN apk add --no-cache ca-certificates

# Copy only the binary from builder
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./server"]

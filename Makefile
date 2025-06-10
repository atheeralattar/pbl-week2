# Makefile for document-system
.PHONY: build run test clean dev docker-up docker-down help

# Build the server binary
build:
	@echo "Building server..."
	@mkdir -p bin
	@go build -o bin/server ./cmd/server

# Run the server
run: build
	@echo "Starting server..."
	@./bin/server

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf bin/

# Start Docker services
docker-up:
	@echo "Starting Docker services..."
	@docker-compose up -d

# Stop Docker services
docker-down:
	@echo "Stopping Docker services..."
	@docker-compose down

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Show help
help:
	@echo "Available commands:"
	@echo "  build      - Build the server binary"
	@echo "  run        - Build and run the server"
	@echo "  dev        - Run server in development mode"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  docker-up  - Start Docker services"
	@echo "  docker-down- Stop Docker services"
	@echo "  deps       - Install and tidy dependencies"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter"
	@echo "  help       - Show this help message" 
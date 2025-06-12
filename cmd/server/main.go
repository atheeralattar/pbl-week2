package main

import (
	"log"
	"os"

	"github.com/atheeralattar/pbl-week2/internal/config"
	"github.com/atheeralattar/pbl-week2/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("[STARTUP] Starting document-system server...")

	// Load environment variables from .env file
	log.Println("[CONFIG] Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Printf("[CONFIG] Warning: Error loading .env file: %v", err)
		log.Println("[CONFIG] Continuing with system environment variables...")
	} else {
		log.Println("[CONFIG] Environment variables loaded successfully")
	}

	// Initialize database
	log.Println("[DATABASE] Initializing database connection...")
	config.InitDB()

	// Setup routes
	log.Println("[ROUTES] Setting up HTTP routes...")
	router := routes.SetupRoutes()

	// Get server port from environment variable
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // fallback to default port
		log.Printf("[CONFIG] SERVER_PORT not set, using default port: %s", port)
	} else {
		log.Printf("[CONFIG] Using configured port: %s", port)
	}

	// Start the server
	log.Printf("[SERVER] Starting HTTP server on port %s...", port)
	log.Printf("[SERVER] Server ready! Visit http://localhost:%s/documents", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("[SERVER] Failed to start server: %v", err)
	}
}

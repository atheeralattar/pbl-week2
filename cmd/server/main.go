package main

import (
	"document-system/internal/config"
	"document-system/internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	config.InitDB()

	// Setup routes
	router := routes.SetupRoutes()

	// Get server port from environment variable
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // fallback to default port
	}

	// Start the server
	router.Run(":" + port)
}

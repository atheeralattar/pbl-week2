package main

import (
	"document-system/config"
	"document-system/routes"
)

func main() {
	// Initialize database
	config.InitDB()

	// Setup routes
	router := routes.SetupRoutes()

	// Start the server
	router.Run(":8080") // localhost:8080
}

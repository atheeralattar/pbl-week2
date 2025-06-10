// routes/routes.go
// This file contains the routes for the document resource
// It defines the SetupRoutes function
// It uses the Gin framework to handle the HTTP requests
// It uses the GORM library to interact with the database

package routes

import (
	"document-system/internal/config"
	"document-system/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Initialize controllers
	documentController := handlers.NewDocumentController(config.GetDB())

	// Document routes
	documentRoutes := r.Group("/documents")
	{
		documentRoutes.POST("", documentController.Create)
		documentRoutes.GET("", documentController.GetAll)
		documentRoutes.GET("/:id", documentController.GetByID)
		documentRoutes.PUT("/:id", documentController.Update)
		documentRoutes.DELETE("/:id", documentController.Delete)
	}

	return r
}

// controllers/document_controller.go
// This file contains the controller for the document resource
// It defines the DocumentController struct and the NewDocumentController function
// It also contains the methods for the DocumentController struct
// It uses the Gin framework to handle the HTTP requests
// It uses the GORM library to interact with the database

package handlers

import (
	"document-system/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentController struct {
	documentModel *models.DocumentModel
}

func NewDocumentController(db *gorm.DB) *DocumentController {
	return &DocumentController{
		documentModel: models.NewDocumentModel(db),
	}
}

func (dc *DocumentController) Create(c *gin.Context) {
	var document models.Document
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the document to the database
	if err := dc.documentModel.Create(&document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the newly created document
	c.JSON(http.StatusCreated, document)
}

func (dc *DocumentController) GetAll(c *gin.Context) {
	documents, err := dc.documentModel.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, documents)
}

func (dc *DocumentController) GetByID(c *gin.Context) {
	id := c.Param("id")
	document, err := dc.documentModel.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	c.JSON(http.StatusOK, document)
}

func (dc *DocumentController) Update(c *gin.Context) {
	id := c.Param("id")
	document, err := dc.documentModel.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	// Bind the incoming data to the existing document
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the document in the database
	if err := dc.documentModel.Update(&document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, document)
}

func (dc *DocumentController) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := dc.documentModel.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	// Delete the document
	if err := dc.documentModel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted"})
}

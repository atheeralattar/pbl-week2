// models/document.go
// This file contains the models for the document resource
// It defines the Document struct and the DocumentModel struct
// It also contains the methods for the DocumentModel struct
// It uses the GORM library to interact with the database

package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID        uint32    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" binding:"required"`
	Content   []byte    `json:"content" binding:"required"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentModel handles all database operations for documents
type DocumentModel struct {
	DB *gorm.DB
}

// NewDocumentModel creates a new document model
func NewDocumentModel(db *gorm.DB) *DocumentModel {
	return &DocumentModel{DB: db}
}

// Create adds a new document to the database
func (m *DocumentModel) Create(doc *Document) error {
	doc.CreatedAt = time.Now()
	doc.UpdatedAt = time.Now()
	return m.DB.Create(doc).Error
}

// FindAll retrieves all documents
func (m *DocumentModel) FindAll() ([]Document, error) {
	var documents []Document
	err := m.DB.Find(&documents).Error
	return documents, err
}

// FindByID retrieves a document by its ID
func (m *DocumentModel) FindByID(id string) (Document, error) {
	var document Document
	err := m.DB.First(&document, id).Error
	return document, err
}

// Update modifies an existing document
func (m *DocumentModel) Update(doc *Document) error {
	doc.UpdatedAt = time.Now()
	return m.DB.Save(doc).Error
}

// Delete removes a document
func (m *DocumentModel) Delete(id string) error {
	return m.DB.Delete(&Document{}, id).Error
}

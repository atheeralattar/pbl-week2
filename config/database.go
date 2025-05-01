// config/database.go
// This file contains the configuration for the database
// It defines the InitDB function
// It uses the GORM library to interact with the database

package config

import (
	"document-system/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=documents_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connected!")

	// Auto migrate the Document model
	err = DB.AutoMigrate(&models.Document{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	fmt.Println("Database Migrated!")
}

func GetDB() *gorm.DB {
	return DB
} 
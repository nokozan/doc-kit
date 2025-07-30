package db

import (
	"doc-kit/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(&models.GitRepo{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("database migrated successfully")
}

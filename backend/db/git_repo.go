package db

import (
	"doc-kit/models"

	"gorm.io/gorm"
)

func GetGitRepoByID(db *gorm.DB, id uint) (models.GitRepo, error) {
	var repo models.GitRepo
	result := db.First(&repo, id)
	return repo, result.Error
}

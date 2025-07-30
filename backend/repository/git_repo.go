package repository

import (
	"doc-kit/models"

	"gorm.io/gorm"
)

type GitRepoRepository struct {
	db *gorm.DB
}

func NewGitRepoRepository(db *gorm.DB) *GitRepoRepository {
	return &GitRepoRepository{db: db}
}

func (r *GitRepoRepository) Add(repo *models.GitRepo) error {
	return r.db.Create(repo).Error
}

func (r *GitRepoRepository) List() ([]models.GitRepo, error) {
	var repos []models.GitRepo
	err := r.db.Order("created_at DESC").Find(&repos).Error
	return repos, err
}

package service

import (
	"doc-kit/db"
	"doc-kit/models"
	"doc-kit/utils/git_utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRepo(ctx *gin.Context) {
	var req models.GitRepo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	repo := &models.GitRepo{
		Alias:       req.Alias,
		URL:         req.URL,
		Branch:      req.Branch,
		Description: req.Description,
	}

	if err := db.DB.Create(repo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save repo", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, repo)
}

func ListRepos(ctx *gin.Context) {
	var repos []models.GitRepo
	if err := db.DB.Order("created_at DESC").Find(&repos).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch repos", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, repos)
}

func CloneRepo(ctx *gin.Context) {
	id := ctx.Param("id")

	var repo models.GitRepo
	if err := db.DB.First(&repo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "repo not found"})
		return
	}

	if err := git_utils.CloneRepo(repo.URL, repo.Branch, repo.Alias); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clone repo", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "repo cloned successfully", "repo": repo.Alias})
}

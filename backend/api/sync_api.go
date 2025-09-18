package api

import (
	"doc-kit/db"
	"doc-kit/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SyncRepoHandler(dbConn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		repo, err := db.GetGitRepoByID(dbConn, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Repo not found"})
			return
		}

		if err := service.SyncRepoStructs(dbConn, repo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Sync failed", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Sync complete"})
	}
}

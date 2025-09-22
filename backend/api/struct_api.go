package api

import (
	"doc-kit/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStructsByRepoID(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repo ID"})
			return
		}

		var structs []models.Struct
		if err := db.Preload("Fields").Preload("Methods").Where("repo_id = ?", id).Find(&structs).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch structs", "details": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, structs)
	}
}

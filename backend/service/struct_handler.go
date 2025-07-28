package service

import (
	"doc-kit/core"
	"doc-kit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var mockStructs = []models.StructMeta{
	{
		Name: "User",
		Doc:  "User struct represents a user in the system.",
		Fields: []models.FieldMeta{
			{Name: "ID", Type: "Unique identifier for the user.", Tag: "`json:\"id\"`", Comment: "Primary key"},
			{Name: "Name", Type: "Name of the user.", Tag: "`json:\"name\"`", Comment: "User's full name"},
		},
		Methods: []models.MethodMeta{
			{Name: "GetFullName", Doc: "Returns the full name of the user."},
			{Name: "IsActive", Doc: "Checks if the user is active."},
		},
	},
}

func GetAllStructs(isDummy bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getStructList := func(
			structs []models.StructMeta,
		) []map[string]string {
			var list []map[string]string
			for _, s := range structs {
				list = append(list, map[string]string{
					"name": s.Name})
			}
			return list
		}
		if isDummy {
			ctx.JSON(http.StatusOK, getStructList(mockStructs))
		} else {
			ctx.JSON(http.StatusOK, getStructList(core.GetAllStructs()))
		}
	}
}

func GetOneStruct(isDummy bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		var meta *models.StructMeta
		if isDummy {
			for i := range mockStructs {
				if mockStructs[i].Name == name {
					meta = &mockStructs[i]
					break
				}
			}
		} else {
			if m, ok := core.GetStructByName(name); ok {
				meta = &m
			}
		}

		if meta == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Struct not found"})
		} else {
			ctx.JSON(http.StatusOK, meta)
		}

	}
}

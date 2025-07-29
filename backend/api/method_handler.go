package api

import (
	"doc-kit/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunMethod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		structName := ctx.Param("struct")
		methodName := ctx.Param("method")

		// Input JSON body -> map
		var input map[string]any
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format", "details": err.Error()})
			return
		}

		// Call core.Run
		result, err := core.Run(structName, methodName, input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return result
		ctx.JSON(http.StatusOK, gin.H{"result": result})

	}
}

package main

import (
	"doc-kit/api"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {

	// db.InitDB()
	// db.Migrate()

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		// ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}
		ctx.Next()
	})

	api.RegisterRoutes(r)

	r.Run(":8080")

}

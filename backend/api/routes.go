package api

import (
	"doc-kit/db"
	"doc-kit/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	// Git repo management routes
	api.POST("/repos", service.CreateRepo)
	api.GET("/repos", service.ListRepos)
	api.POST("/repos/:id/clone", service.CloneRepo)

	api.GET("/structs")
	api.GET("/structs", service.GetAllStructs(false))
	api.GET("/structs_dummy", service.GetAllStructs(true))
	// r.GET("/structs/:name", service.GetOneStruct)
	// r.POST("/run/:struct/:method", runMethod)
	api.POST("/repos/:id/sync", SyncRepoHandler(db.DB))
	api.GET("/repos/:id/structs", GetStructsByRepoID(db.DB))
}

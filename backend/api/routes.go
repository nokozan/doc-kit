package api

import (
	"doc-kit/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/structs")
	r.GET("/structs", service.GetAllStructs(false))
	r.GET("/structs_dummy", service.GetAllStructs(true))
	// r.GET("/structs/:name", service.GetOneStruct)
	// r.POST("/run/:struct/:method", runMethod)

}

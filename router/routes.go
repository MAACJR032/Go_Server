package router

import (
	"API/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.CreateOpening)
		v1.POST("/opening", handler.ShowOpening)
		v1.DELETE("/opening", handler.DeleteOpening)
		v1.PUT("/opening", handler.UpdateOpening)
		v1.GET("/openings", handler.ListOpenings)
	}
}
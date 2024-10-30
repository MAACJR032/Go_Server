package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initRoutes(router)

	router.Run(":8080")
}
package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api/v1")

	AccountRoutes(apiGroup)

	return router
}

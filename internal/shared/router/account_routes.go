package router

import (
	config "trilha-api/internal/shared/config"
	"trilha-api/internal/wire"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(apiGroup *gin.RouterGroup) {
	accountHandler := wire.NewAccountHandler(config.DB)

	accountGroup := apiGroup.Group("/accounts")

	accountGroup.POST("/", accountHandler.Register)
}

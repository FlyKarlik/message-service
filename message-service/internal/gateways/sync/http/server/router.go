package server

import (
	_ "github.com/FlyKarlik/message-service/api/docs"
	"github.com/FlyKarlik/message-service/internal/gateways/sync/http/synchandle"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(syncHandler *synchandle.SyncHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-type"},
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api")
	api.Use(syncHandler.SetJSON)
	{
		messageGroup := api.Group("message")
		{
			messageGroup.GET("/", syncHandler.GetAllMessage)
			messageGroup.GET("/processed", syncHandler.GetAllProcessedMessage)
			messageGroup.POST("/", syncHandler.AddMessage)
			messageGroup.GET("/:id", syncHandler.GetMessage)

			statsGroup := messageGroup.Group("stats")
			{
				statsGroup.GET("/", syncHandler.GetStats)
			}

		}

	}

	return router
}

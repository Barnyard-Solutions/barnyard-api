package routers

import (
	"barnyard/api/handlers"
	"barnyard/api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	router.GET("/", handlers.GetDefaultMessage)
	v1 := router.Group("/api/v1")

	userRoutes := v1.Group("/user")
	{
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.POST("/token", handlers.GetToken)
	}

	feedRoutes := v1.Group("/feed")
	{
		feedRoutes.Use(middlewares.TokenAuthMiddleware())
		feedRoutes.GET("/", handlers.GetFeeds)
		feedRoutes.POST("/", handlers.CreateFeed)
		feedRoutes.DELETE("/", handlers.RemoveFeed)
		feedRoutes.POST("/event", handlers.CreateEvent)
		feedRoutes.GET("/:id/event", handlers.GetEvent)
	}

}

package routers

import (
	"barnyard/api/handlers"

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

	feedsRoutes := v1.Group("/feed")
	{
		feedsRoutes.GET("/", handlers.GetFeeds)
	}

}

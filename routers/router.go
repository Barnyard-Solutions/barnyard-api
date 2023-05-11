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
		feedRoutes.DELETE("/:id", handlers.RemoveFeed)

		eventRoutes := feedRoutes.Group("/:id/event")
		{
			eventRoutes.POST("/", handlers.CreateEvent)
			eventRoutes.GET("/", handlers.GetEvent)
			eventRoutes.DELETE("/:eventID", handlers.RemoveEvent)
		}

		milestoneRoutes := feedRoutes.Group("/:id/milestone")
		{
			milestoneRoutes.POST("/", handlers.CreateMilestone)
			milestoneRoutes.GET("/", handlers.GetMilestone)
			///milestoneRoutes.DELETE("/:eventID", handlers.RemoveEvent)
		}

	}

}

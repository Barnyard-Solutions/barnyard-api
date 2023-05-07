package routers

import (
	"barnyard/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	router.GET("/", handlers.GetDefaultMessage)
	v1 := router.Group("/api/v1")

	// User routes
	userRoutes := v1.Group("/user")
	{
		//userRoutes.GET("/:id", handlers.GetUser)
		userRoutes.POST("/", handlers.CreateUser)
		//userRoutes.PUT("/:id", handlers.UpdateUser)
		//userRoutes.DELETE("/:id", handlers.DeleteUser)
	}

	// Add more routes for other resources as needed
}

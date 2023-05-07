package main

import (
	"github.com/gin-gonic/gin"

	"barnyard/api/routers"
)

// func getDefaultMessage(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "👋 Barnyard API is up <br> ")
// }

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.svg")
	})

	routers.SetUp(router)

	router.Run("localhost:5000")
}

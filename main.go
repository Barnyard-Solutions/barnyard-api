package main

import (
	"github.com/gin-gonic/gin"

	"barnyard/api/database"
	"barnyard/api/routers"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	err := database.InitDB("API:kKQdH@qX93@tcp(localhost:3306)/barnyard")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer database.CloseDB()

	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.svg")
	})

	routers.SetUp(router)

	router.Run("localhost:5000")

}

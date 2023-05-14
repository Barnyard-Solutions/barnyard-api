package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"barnyard/api/database"
	"barnyard/api/routers"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbURL := fmt.Sprintf("API:kKQdH@qX93@tcp(%s:3306)/barnyard?timeout=120s", dbHost)

	err := database.InitDB(dbURL)
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

	router.Run("0.0.0.0:5000")

}

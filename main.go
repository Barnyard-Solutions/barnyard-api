package main

import (
	"github.com/gin-gonic/gin"

	"barnyard/api/database"
	"barnyard/api/routers"

	"database/sql"
	"fmt"
	"log"

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

func testDB() {
	db, err := sql.Open("mysql", "API:kKQdH@qX93@tcp(localhost:3306)/barnyard")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Perform the query
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process the query results
	fmt.Println("results:")
	for rows.Next() {
		var id int
		var name string
		var pass string
		err := rows.Scan(&id, &name, &pass)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID:", id, "Name:", name, "Pass:", pass)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

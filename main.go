package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func getDefaultMessage(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "👋 Barnyard API is up <br> ")
// }

func getDefaultMessage(c *gin.Context) {
	//c.Header("Content-Type", "text/html")
	//c.String(http.StatusOK, html.EscapeString("👋")+)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("👋 Barnyard API is up <br> check github repo <a href='https://github.com/Barnyard-Solutions/barnyard-api' >here</a>"))
}

func main() {
	router := gin.Default()
	router.GET("/", getDefaultMessage)

	router.Run("localhost:5000")
}

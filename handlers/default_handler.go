package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDefaultMessage(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("👋 Barnyard API is up <br> check github repo <a href='https://github.com/Barnyard-Solutions/barnyard-api' >here</a>"))
}

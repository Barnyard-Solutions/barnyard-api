package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"

	"github.com/gin-gonic/gin"
)

func GetFeeds(c *gin.Context) {
	token := c.Query("token")
	var feeds []models.Feed

	if token == "" {
		c.JSON(400, gin.H{"error": "Token is required"})
		return
	}

	feeds = database.GetFeeds(token)
	if feeds == nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve user's feeds"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User feeds retrieve successfully",
		"feeds":   feeds,
	})
}

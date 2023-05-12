package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSubscription(c *gin.Context) {
	var requestBody models.CreateSubscriptionRequest
	token := c.GetString("token")
	feedID := c.Param("id")

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	success, err := database.InsertSubscription(token, requestBody.Subscription, feedIDInt)
	if err != nil || !success {
		c.JSON(500, gin.H{"error": "Failed to create user's subscription"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User subscription created successfully",
	})
}

/*

func GetSubscriptions(c *gin.Context) {
	token := c.GetString("token")
	feeds, err := database.SelectFeeds(token)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve user's feeds"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User feeds retrieved successfully",
		"feeds":   feeds,
	})
}

func RemoveFeed(c *gin.Context) {
	token := c.GetString("token")
	feedID := c.Param("id")

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	success, err := database.DeleteFeed(token, feedIDInt)
	if err != nil || !success {
		c.JSON(500, gin.H{"error": err})

		return
	}

	c.JSON(200, gin.H{
		"message": "User feed removed successfully",
	})

}
*/

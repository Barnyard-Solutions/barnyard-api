package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSubscription(c *gin.Context) {
	var requestBody models.SubscriptionRequest
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

func IsSubscribed(c *gin.Context) {
	token := c.GetString("token")
	feedID := c.Param("id")
	endPoint := c.Param("endPoint")
	result := false

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	subscriptions, err := database.SelectSubscription(token, endPoint, feedIDInt)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve user's subscriptions"})
		return
	}

	if len(subscriptions) >= 1 {
		result = true
	}

	c.JSON(200, gin.H{
		"message":    "User subscriptions retrieved successfully",
		"subscribed": result,
	})
}

func RemoveSubscription(c *gin.Context) {
	var requestBody models.SubscriptionRequest
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

	success, err := database.DeleteSubscription(token, requestBody.Subscription, feedIDInt)
	if err != nil || !success {
		if err == database.ErrUnauthorized {
			c.JSON(401, gin.H{"error": "Unauthorized"})
		} else if err == database.ErrNotFound {
			c.JSON(404, gin.H{"error": "Subscription not found"})
		} else {
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "User feed removed successfully",
	})

}

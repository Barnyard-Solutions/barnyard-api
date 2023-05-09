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

	feeds = database.SelectFeeds(token)
	if feeds == nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve user's feeds"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User feeds retrieve successfully",
		"feeds":   feeds,
	})
}

func CreateFeed(c *gin.Context) {
	var requestBody models.CreateFeedRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if requestBody.Token == "" {
		c.JSON(400, gin.H{"error": "Token is required"})
		return
	}

	success, err := database.InsertFeed(requestBody.Token, requestBody.FeedName)
	if err != nil || !success {
		if database.IsDuplicateEntryError(err) {
			c.JSON(400, gin.H{"error": "Feed name already exists"})
		} else {
			c.JSON(500, gin.H{"error": "Failed to create user's feed"})
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "User feed created successfully",
	})
}

func RemoveFeed(c *gin.Context) {
	var requestBody models.DeleteFeedRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if requestBody.Token == "" {
		c.JSON(400, gin.H{"error": "Token is required"})
		return
	}

	success, err := database.DeleteFeed(requestBody.Token, requestBody.FeedID)
	if err != nil || !success {
		c.JSON(500, gin.H{"error": err})

		return
	}

	c.JSON(200, gin.H{
		"message": "User feed created successfully",
	})

}

func CreateEvent(c *gin.Context) {
	var requestBody models.CreateEventRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if requestBody.Token == "" {
		c.JSON(400, gin.H{"error": "Token is required"})
		return
	}

	success, err := database.InsertEvent(requestBody.Token, requestBody.Name1,
		requestBody.Name2, requestBody.Date, requestBody.FeedID)
	if err != nil || !success {
		c.JSON(500, gin.H{"error": err})

		return
	}

	c.JSON(200, gin.H{
		"message": "Event created successfully",
	})

}

func GetEvent(c *gin.Context) {
	var requestBody models.GetEventRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if requestBody.Token == "" {
		c.JSON(400, gin.H{"error": "Token is required"})
		return
	}

	events, err := database.SelectEvents(requestBody.Token, requestBody.FeedID)
	if err != nil {
		c.JSON(500, gin.H{"error": err})

		return
	}

	c.JSON(200, gin.H{
		"message": "Event created successfully",
		"events":  events,
	})

}

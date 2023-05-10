package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFeeds(c *gin.Context) {
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

func CreateFeed(c *gin.Context) {
	var requestBody models.CreateFeedRequest
	token := c.GetString("token")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	success, err := database.InsertFeed(token, requestBody.FeedName)
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
	token := c.GetString("token")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	success, err := database.DeleteFeed(token, requestBody.FeedID)
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
	token := c.GetString("token")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	success, err := database.InsertEvent(token, requestBody.Name1,
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
	feedID := c.Param("id")
	token := c.GetString("token")

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	events, err := database.SelectEvents(token, feedIDInt)
	if err != nil {
		c.JSON(500, gin.H{"error": err})

		return
	}

	c.JSON(200, gin.H{
		"message": "Event created successfully",
		"events":  events,
	})

}

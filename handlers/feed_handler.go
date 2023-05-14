package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFeeds(c *gin.Context) {
	token := c.GetString("token")
	endPoint := c.Query("endpoint") // Extract the value of "endpoint" query parameter

	if endPoint != "" {
		feeds, err := database.SelectFeedsWithSub(token, endPoint)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve user's feeds"})
			return
		}

		c.JSON(200, gin.H{
			"message": "User feeds retrieved successfully",
			"feeds":   feeds,
		})

	} else {
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
		if err == database.ErrUnauthorized {
			c.JSON(401, gin.H{"error": "Unauthorized"})
		} else if err == database.ErrNotFound {
			c.JSON(404, gin.H{"error": "Feed not found"})
		} else {
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "User feed removed successfully",
	})

}

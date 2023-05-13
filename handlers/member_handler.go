package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMember(c *gin.Context) {
	var requestBody models.MemberRequest
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

	success, err := database.InsertMemeber(token, requestBody.Mail, feedIDInt, requestBody.Permission)
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
		"message": "Feed member created successfully",
	})
}

func GetMembers(c *gin.Context) {
	token := c.GetString("token")
	feedID := c.Param("id")

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	members, err := database.SelectMembers(token, feedIDInt)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve user's feeds"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Feed members retrieved successfully",
		"members": members,
	})
}

/*

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
*/

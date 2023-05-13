package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMilestone(c *gin.Context) {
	var requestBody models.CreateMilestoneRequest
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

	success, err := database.InsertMilestone(token, requestBody.Name,
		requestBody.Color, feedIDInt, requestBody.Date)
	if err != nil || !success {
		c.JSON(500, gin.H{"error": err})

		return
	}

	c.JSON(200, gin.H{
		"message": "Milestone created successfully",
	})

}

func GetMilestone(c *gin.Context) {
	feedID := c.Param("id")
	token := c.GetString("token")

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	milestones, err := database.SelectMilestone(token, feedIDInt)
	if err != nil {
		c.JSON(500, gin.H{"error": err})

		return
	}

	c.JSON(200, gin.H{
		"message":    "Milestones retrieved successfully",
		"milestones": milestones,
	})

}

func RemoveMilestone(c *gin.Context) {
	token := c.GetString("token")
	feedID := c.Param("id")
	milestoneID := c.Param("milestoneID")

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	// Convert milestoneID to an integer
	milestoneIDInt, err := strconv.Atoi(milestoneID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid milestone ID"})
		return
	}

	success, err := database.DeleteMilestone(token, feedIDInt, milestoneIDInt)
	if err != nil || !success {
		if err == database.ErrUnauthorized {
			c.JSON(401, gin.H{"error": "Unauthorized"})
		} else if err == database.ErrNotFound {
			c.JSON(404, gin.H{"error": "Milestone not found"})
		} else {
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "User milestone removed successfully",
	})

}

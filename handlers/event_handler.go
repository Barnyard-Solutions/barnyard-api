package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var requestBody models.CreateEventRequest
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

	success, err := database.InsertEvent(token, requestBody.Name1,
		requestBody.Name2, requestBody.Date, feedIDInt)
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
		"message": "Events retrieved successfully",
		"events":  events,
	})

}

func RemoveEvent(c *gin.Context) {
	token := c.GetString("token")
	feedID := c.Param("id")
	eventID := c.Param("eventID")

	// Convert feedID to an integer
	feedIDInt, err := strconv.Atoi(feedID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid feed ID"})
		return
	}

	// Convert eventID to an integer
	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	success, err := database.DeleteEvent(token, feedIDInt, eventIDInt)
	if err != nil || !success {
		if err == database.ErrUnauthorized {
			c.JSON(401, gin.H{"error": "Unauthorized"})
		} else if err == database.ErrNotFound {
			c.JSON(404, gin.H{"error": "Event not found"})
		} else {
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "User event removed successfully",
	})

}

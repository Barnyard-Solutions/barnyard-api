package handlers

import (
	"barnyard/api/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.PassKey == "" {
		c.JSON(400, gin.H{"error": "User Email and PassKey are required"})
		return
	}

	// Access the values from the s struct
	//userID := user.ID
	//userName := user.Name

	// Perform desired operations with the retrieved values

	c.JSON(200, gin.H{
		"message": "Event created",
	})
}

package handlers

import (
	"barnyard/api/database"
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

	err := database.InsertUser(user.Email, user.PassKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to insert user into the database"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

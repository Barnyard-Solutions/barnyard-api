package handlers

import (
	"barnyard/api/database"
	"barnyard/api/models"
	"fmt"

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

func GetUser(c *gin.Context) {
	token := c.GetString("token")

	user, err := database.SelectUser(token)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User selected successfully",
		"user":    user,
	})
}

func GetToken(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.PassKey == "" {
		c.JSON(400, gin.H{"error": "User Email and PassKey are required"})
		return
	}

	token, err := database.GenerateUserToken(user.Email, user.PassKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token into the database"})
		fmt.Println("Error:", err)
		return
	}

	if token == "" {
		c.JSON(400, gin.H{"error": "User Email or PassKey are wrong"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Token generate successfully",
		"token":   token,
	})
}

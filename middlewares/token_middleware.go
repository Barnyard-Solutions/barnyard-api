package middlewares

import (
	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware is a middleware function to extract and validate the session token
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(400, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// Remove any token type prefix if present
		token = removeTokenTypePrefix(token)

		c.Set("token", token)
		c.Next()
	}
}

// Function to remove token type prefix if present
func removeTokenTypePrefix(token string) string {
	// Check if token starts with "Bearer " prefix
	if len(token) > 7 && token[:7] == "Bearer " {
		return token[7:]
	}
	return token
}

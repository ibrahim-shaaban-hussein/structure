package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function to authenticate requests using Bearer Tokens
func AuthMiddleware() gin.HandlerFunc {
	// Define the middleware function that Gin will execute for each request
	return func(c *gin.Context) {
		// Get the Authorization header from the request
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is missing or doesn't start with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			// If the header is missing or incorrect, return a 401 Unauthorized response
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Stop processing the request further
			return
		}

		// Extract the token from the Authorization header
		tokenString := authHeader[len("Bearer "):]

		// Validate the token
		if !isValidToken(tokenString) {
			// If the token is invalid, return a 401 Unauthorized response
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() // Stop processing the request further
			return
		}

		// Proceed to the next middleware or handler
		c.Next()
	}
}

// isValidToken is a function to validate the token
func isValidToken(token string) bool {
	// Example token validation logic (replace with your actual implementation)
	// For demonstration purposes, we'll assume the token is valid if its length is greater than 10 characters
	if len(token) > 10 {
		return true
	}
	return false
}


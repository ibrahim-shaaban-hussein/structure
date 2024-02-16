// middleware/auth.go

package tests

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
		token := authHeader[len("Bearer "):]

		// Perform token validation logic here, such as verifying token authenticity, expiration, etc.
		// For simplicity, we'll assume token validation is successful for demonstration purposes.

		// Proceed to the next middleware or handler
		c.Next()
	}
}

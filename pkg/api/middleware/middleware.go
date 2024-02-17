package middleware

import (
    "strings"
    "net/http"
    "github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function to authenticate requests using Bearer Tokens
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the Authorization header from the request
        authHeader := c.GetHeader("Authorization")

        // Check if the Authorization header is missing or doesn't start with "Bearer "
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort() // Stop processing the request further
            return
        }

        // Proceed to the next middleware or handler
        c.Next()
    }
}


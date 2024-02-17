package middleware

import (
    "github.com/gin-gonic/gin"
)

// Define a secret key for signing and verifying tokens
var secretKey = []byte("your-secret-key")

// ExampleAuthMiddleware is an example middleware specific to authentication
func ExampleAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Your authentication logic here
        c.Next() // Continue processing
    }
}


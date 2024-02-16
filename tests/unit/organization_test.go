package tests

import (
    "fmt" // Import the fmt package
    "net/http"
    "strings"
    "time"

    "github.com/dgrijalva/jwt-go" // Import the JWT library
    "github.com/gin-gonic/gin"
)

// Define a secret key for signing and verifying tokens
var secretKey = []byte("your-secret-key")

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

        // Parse the token string into a JWT token object
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Ensure the token signing method is HMAC and the signing key matches
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })
        if err != nil {
            // If there's an error parsing the token, return a 401 Unauthorized response
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort() // Stop processing the request further
            return
        }

        // Check if the token is valid
        if !token.Valid {
            // If the token is invalid, return a 401 Unauthorized response
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort() // Stop processing the request further
            return
        }

        // Check token expiration
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            // Extract and validate the "exp" claim to check token expiration
            expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
            if time.Now().After(expirationTime) {
                // If the token is expired, return a 401 Unauthorized response
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
                c.Abort() // Stop processing the request further
                return
            }
        }

        // Proceed to the next middleware or handler
        c.Next()
    }
}

// Write your test functions here...


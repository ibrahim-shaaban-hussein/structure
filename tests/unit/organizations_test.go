package tests

import (
    "net/http"

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

// UpdateOrganizationHandler handles the request for updating an organization
func UpdateOrganizationHandler(c *gin.Context) {
    // Extract organization ID from the request parameters
    organizationID := c.Param("organization_id")

    // Parse the request body to extract the updated organization data
    var requestBody struct {
        Name        string `json:"name"`
        Description string `json:"description"`
    }
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Perform update logic
    // Here, you would typically update the organization in your database or any other data store
    // For demonstration purposes, let's assume we have a function named UpdateOrganizationByID to update the organization by its ID
    err := UpdateOrganizationByID(organizationID, requestBody.Name, requestBody.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update organization"})
        return
    }

    // Return the updated organization details as JSON response
    updatedOrganization := map[string]interface{}{
        "organization_id": organizationID,
        "name":            requestBody.Name,
        "description":     requestBody.Description,
    }
    c.JSON(http.StatusOK, updatedOrganization)
}

// DeleteOrganizationHandler handles the request for deleting an organization
func DeleteOrganizationHandler(c *gin.Context) {
    // Extract organization ID from the request parameters
    organizationID := c.Param("organization_id")

    // Perform delete logic
    // Here, you would typically delete the organization from your database or any other data store
    // For demonstration purposes, let's assume we have a function named DeleteOrganizationByID to delete the organization by its ID
    err := DeleteOrganizationByID(organizationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organization"})
        return
    }

    // Return success message as JSON response
    c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}

// InviteUserToOrganizationHandler handles the request for inviting a user to an organization
func InviteUserToOrganizationHandler(c *gin.Context) {
    // Extract organization ID from the request parameters
    organizationID := c.Param("organization_id")

    // Extract user email from the request body
    var requestBody struct {
        UserEmail string `json:"user_email"`
    }
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Perform invitation logic
    // Here, you would typically send an invitation to the user to join the organization
    // For demonstration purposes, let's assume we have a function named InviteUserToOrganization to invite the user to the organization
    err := InviteUserToOrganization(organizationID, requestBody.UserEmail)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invite user to organization"})
        return
    }

    // Return success message as JSON response
    c.JSON(http.StatusOK, gin.H{"message": "User invited to organization successfully"})
}


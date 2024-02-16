// handlers/invite_user.go

package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// InviteUserRequest represents the JSON request body schema for inviting a user to an organization
type InviteUserRequest struct {
    UserEmail string `json:"user_email"`
}

// InviteUserHandler handles the invitation of a user to an organization
func InviteUserHandler(c *gin.Context) {
    // Extract organization ID from the URL path parameters
    organizationID := c.Param("organization_id")

    // Parse request body
    var req InviteUserRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Perform the invitation logic (e.g., send invitation email, update database)

    // Respond with success message
    c.JSON(http.StatusOK, gin.H{"message": "User invited to organization successfully", "organization_id": organizationID, "user_email": req.UserEmail})
}


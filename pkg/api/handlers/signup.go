package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Define a struct representing the expected JSON format of the request body for signup
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handler function for the /signup endpoint
func SignupHandler(c *gin.Context) {
	// Initialize an empty struct to hold the parsed JSON data
	var req SignupRequest

	// Parse the request body and bind it to the struct
	if err := c.BindJSON(&req); err != nil {
		// If there's an error parsing the JSON, return a bad request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Access the fields of the struct to retrieve the required information
	name := req.Name
	email := req.Email

	// Perform further processing with the extracted information
	// For example, you can validate the input, create a user, etc.

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully", "name": name, "email": email})
}


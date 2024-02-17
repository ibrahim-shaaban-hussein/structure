package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"
	"github.com/ibrahim-shaaban-hussein/structure/pkg/api/middleware"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Register handler functions for different endpoints
	router.GET("/", handlers.HomeHandler)
	router.POST("/refresh-token", handlers.RefreshTokenHandler)
	router.POST("/organizations", handlers.CreateOrganizationHandler)
	router.GET("/organizations/:id", handlers.ReadOrganizationHandler)      // Endpoint for reading a specific organization
	router.GET("/organizations", handlers.ReadAllOrganizationsHandler)     // Endpoint for reading all organizations
	router.POST("/organizations/:id/invite", handlers.InviteUserHandler)    // Invite User to Organization Endpoint

	// Start the HTTP server
	port := ":8080" // specify the port to listen on
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := router.Run(port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}


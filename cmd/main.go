package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Register handler functions for different endpoints
	router.GET("/", handlers.HomeHandler)
	router.POST("/refresh-token", handlers.RefreshTokenHandler)
	// Implement CreateOrganizationHandler and ReadOrganizationHandler
	// Example:
	// router.POST("/organizations", handlers.CreateOrganizationHandler)
	// router.GET("/organizations/:id", handlers.ReadOrganizationHandler)
	router.POST("/organizations/:id/invite", handlers.InviteUserHandler)

	// Register the /organization endpoint
	router.GET("/organization", handlers.OrganizationHandler)

	// Start the HTTP server
	port := ":8080" // specify the port to listen on
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := router.Run(port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}


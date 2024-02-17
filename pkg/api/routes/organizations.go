package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupOrganizationsRoutes registers routes related to managing multiple organizations
func SetupOrganizationsRoutes(router *gin.Engine) {
	router.GET("/organizations", handlers.ReadAllOrganizationsHandler)
}

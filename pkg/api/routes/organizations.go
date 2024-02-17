package routes

import (
	"your-project/handlers"

	"github.com/gin-gonic/gin"
	"github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"
)

// SetupOrganizationsRoutes registers routes related to managing multiple organizations
func SetupOrganizationsRoutes(router *gin.Engine) {
	router.GET("/organizations", handlers.ReadAllOrganizationsHandler)
}

package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupOrganizationRoutes registers routes related to organization management
func SetupOrganizationRoutes(router *gin.Engine) {
	router.PUT("/organization/:organization_id", handlers.UpdateOrganizationHandler)
}

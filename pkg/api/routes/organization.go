package routes

import (
	"github.com/gin-gonic/gin"
	//"your-project/handlers"
	"github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"
)

// SetupOrganizationRoutes registers routes related to organization management
func SetupOrganizationRoutes(router *gin.Engine) {
	router.PUT("/organization/:organization_id", handlers.UpdateOrganizationHandler)
}

// routes/refresh_token.go

package routes

import (
	"github.com/gin-gonic/gin"
        "github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"
)

// SetupRefreshTokenRoutes registers routes related to refreshing tokens
func SetupRefreshTokenRoutes(router *gin.Engine) {
	router.POST("/refresh-token", handlers.RefreshTokenHandler)
}

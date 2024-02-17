// routes/signin.go

package routes

import (
	"IDEANEST/handlers"

	"github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupSigninRoutes registers routes related to user signin
func SetupSigninRoutes(router *gin.Engine) {
	router.POST("/signin", handlers.SigninHandler)
}

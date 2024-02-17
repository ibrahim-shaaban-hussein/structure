// routes/signin.go

package routes

import (
	"github.com/gin-gonic/gin"
        "github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"
)

// SetupSigninRoutes registers routes related to user signin
func SetupSigninRoutes(router *gin.Engine) {
	router.POST("/signin", handlers.SigninHandler)
}

// routes/signup.go

package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupSignupRoutes registers routes related to user signup
func SetupSignupRoutes(router *gin.Engine) {
	router.POST("/signup", handlers.SignupHandler)
}

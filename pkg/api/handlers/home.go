package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandler handles requests to the home page
func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, IDEANEST!, by Ibrahim")
}


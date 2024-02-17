package tests

import (
        "net/http"
        "net/http/httptest"
        "testing"

        "github.com/gin-gonic/gin"
)

// TestOrganizationHandler is a unit test for the organization handler
func TestOrganizationHandler(t *testing.T) {
        // Create a new Gin router
        router := gin.New()

        // Add the AuthMiddleware to the router
        router.Use(AuthMiddleware())

        // Define a handler function for the organization endpoint
        router.GET("/organization", func(c *gin.Context) {
                // Return a dummy response
                c.JSON(http.StatusOK, gin.H{"message": "Organization endpoint"})
        })

        // Create a new HTTP request to test the organization endpoint
        req, err := http.NewRequest("GET", "/organization", nil)
        if err != nil {
                t.Fatalf("failed to create request: %v", err)
        }

        // Create a response recorder to capture the response
        rr := httptest.NewRecorder()

        // Dispatch the request to the router
        router.ServeHTTP(rr, req)

        // Check if the status code is OK (200)
        if status := rr.Code; status != http.StatusOK {
                t.Errorf("handler returned wrong status code: got %v want %v",
                        status, http.StatusOK)
        }

        // Check the response body
        expected := `{"message":"Organization endpoint"}`
        if rr.Body.String() != expected {
                t.Errorf("handler returned unexpected body: got %v want %v",
                        rr.Body.String(), expected)
        }
}


package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRefreshTokenHandler(t *testing.T) {
	// Create a new Gin router
	r := gin.New()

	// Route for the refresh token handler
	r.POST("/refresh-token", RefreshTokenHandler)

	// Create a test request with a JSON body
	reqBody := `{"refresh_token": "dummy-refresh-token"}`
	req, err := http.NewRequest("POST", "/refresh-token", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expectedResponseBody := `{"message":"Token refreshed successfully","access_token":"dummy-access-token","refresh_token":"dummy-refresh-token"}`
	if w.Body.String() != expectedResponseBody {
		t.Errorf("Unexpected response body: got %s, want %s", w.Body.String(), expectedResponseBody)
	}
}

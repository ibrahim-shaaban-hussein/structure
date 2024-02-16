// signup_test.go

package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/ibrahim-shaaban-hussein/structure/pkg/api/handlers"
	handlers "github.com/ibrahim-shaaban-hussein/structure/pkg/api/middleware"
)

func TestSignupHandler(t *testing.T) {
	// Define the request body
	requestBody := map[string]string{
		"name":     "John Doe",
		"email":    "john@example.com",
		"password": "password123",
	}

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("error marshaling request body: %v", err)
	}

	// Create a request to the Signup endpoint
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the SignupHandler
	http.HandlerFunc(handlers.SignupHandler).ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expectedResponseBody := `{"message":"User signed up successfully","name":"John Doe","email":"john@example.com"}`
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponseBody)
	}
}

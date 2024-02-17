package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

i   "github.com/username/project/handlers"
)

func TestInviteUserHandler(t *testing.T) {
    // Define the request payload
    payload := `{"user_email": "user@example.com"}`

    // Create a request to the invite user endpoint
    req, err := http.NewRequest("POST", "/organization/organization_id/invite", strings.NewReader(payload))
    if err != nil {
        t.Fatalf("error creating request: %v", err)
    }

    // Add authorization header
    req.Header.Set("Authorization", "Bearer example_access_token")

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Call the InviteUserHandler
    http.HandlerFunc(handlers.InviteUserToOrganizationHandler).ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body
    var responseBody map[string]string
    if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
        t.Fatalf("error unmarshalling response body: %v", err)
    }

    // Check if the response message is correct
    expectedMessage := "User invited to organization successfully"
    if responseBody["message"] != expectedMessage {
        t.Errorf("unexpected response message: got %s want %s",
            responseBody["message"], expectedMessage)
    }
}


// e2e_handlers_test.go

package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "testing"

    "github.com/ibrahim-shaaban-hussein/structure/tests/e2e/helpers"
)

func TestEndToEnd(t *testing.T) {
    // Setup
    client := &http.Client{}
    baseURL := "http://localhost:8080" // Update with your API base URL

    // Test scenarios
    t.Run("TestInviteUser", func(t *testing.T) {
        // Prepare request payload
        requestData := map[string]string{
            "user_email": "user@example.com",
        }
        requestDataBytes, _ := json.Marshal(requestData)

        // Send request
        response, err := client.Post(baseURL+"/organization/organization_id/invite", "application/json", bytes.NewBuffer(requestDataBytes))
        if err != nil {
            t.Fatalf("error sending request: %v", err)
        }
        defer response.Body.Close()

        // Verify response status code
        if response.StatusCode != http.StatusOK {
            t.Errorf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
        }

        // Verify response body
        var responseBody map[string]interface{}
        if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
            t.Fatalf("error decoding response body: %v", err)
        }

        // Add more assertions as needed
    })

    // Add more test scenarios for other endpoints
}


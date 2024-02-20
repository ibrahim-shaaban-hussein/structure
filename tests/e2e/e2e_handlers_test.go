// e2e_handlers_test.go

package tests

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "testing"
)

func TestEndToEnd(t *testing.T) {
    // Setup
    client := &http.Client{}
    baseURL := "http://localhost:8080" // Update with your API base URL
    organizationID := "organization_id" // Update with the actual organization ID

    // Test scenarios
    t.Run("TestInviteUser", func(t *testing.T) {
        // Prepare request payload
        requestData := map[string]string{
            "user_email": "user@example.com",
        }
        requestDataBytes, _ := json.Marshal(requestData)

        // Send request
        response, err := client.Post(baseURL+"/organization/"+organizationID+"/invite", "application/json", bytes.NewBuffer(requestDataBytes))
        if err != nil {
            t.Fatalf("error sending request: %v", err)
        }
        defer response.Body.Close()

        // Print response status code for debugging
        fmt.Println("Response Status Code:", response.StatusCode)

        // Verify response status code
        expectedStatusCode := http.StatusOK // Update with the expected status code
        if response.StatusCode != expectedStatusCode {
            t.Errorf("expected status code %d, got %d", expectedStatusCode, response.StatusCode)
        }

        // Verify response body
        var responseBody map[string]interface{}
        if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
            t.Fatalf("error decoding response body: %v", err)
        }

        // Print response body for debugging
        fmt.Println("Response Body:", responseBody)

        // Check if the response contains an error message
        if errorMessage, ok := responseBody["error"].(string); ok {
            t.Errorf("received error from server: %s", errorMessage)
        }

        // Add more assertions as needed
    })

    // Add more test scenarios for other endpoints
}


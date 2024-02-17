package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/username/project/handlers"
)

func TestSigninHandler(t *testing.T) {
    // Define the request payload
    payload := `{"email": "user@example.com", "password": "password"}`

    // Create a request to the signin endpoint
    req, err := http.NewRequest("POST", "/signin", strings.NewReader(payload))
    if err != nil {
        t.Fatalf("error creating request: %v", err)
    }

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Call the SigninHandler
    http.HandlerFunc(handlers.SigninHandler).ServeHTTP(rr, req)

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
    expectedMessage := "Signin successful"
    if responseBody["message"] != expectedMessage {
        t.Errorf("unexpected response message: got %s want %s",
            responseBody["message"], expectedMessage)
    }

    // Check if the access token is present
    if _, ok := responseBody["access_token"]; !ok {
        t.Errorf("access token not found in response")
    }
}


// e2e_helpers.go

package helpers

import (
    "encoding/json"
    "net/http"
    "github.com/your-username/your-project/pkg/api/handlers"
    "github.com/your-username/your-project/tests/e2e/helpers"
)

// SendHTTPRequest sends an HTTP request and returns the response.
func SendHTTPRequest(method, url string, payload interface{}) (*http.Response, error) {
    client := &http.Client{}

    // Prepare request body
    var requestBody []byte
    if payload != nil {
        requestBody, _ = json.Marshal(payload)
    }

    // Create request
    req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }

    // Set request headers and content type
    req.Header.Set("Content-Type", "application/json")

    // Send request
    return client.Do(req)
}

// DecodeJSONResponse decodes the JSON response body into the provided interface.
func DecodeJSONResponse(response *http.Response, target interface{}) error {
    defer response.Body.Close()

    // Decode response body
    return json.NewDecoder(response.Body).Decode(target)
}


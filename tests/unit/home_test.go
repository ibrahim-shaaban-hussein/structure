package tests

import (
    "testing"
    "github.com/your-username/structure/pkg/database/mongodb"
)

// SetupMockDatabase initializes a mock database connection for testing
func SetupMockDatabase() (*MockDBClient, error) {
    // Create a new instance of the mock database client
    db := &MockDBClient{}

    // Connect to the mock database
    if err := db.Connect(); err != nil {
        return nil, err
    }

    // Return the initialized mock database client
    return db, nil
}

// TeardownMockDatabase cleans up the mock database connection
func TeardownMockDatabase(db *MockDBClient) {
    // Disconnect from the mock database
    _ = db.Disconnect()
}

func TestHomeHandler(t *testing.T) {
    // Setup test environment
    mockDB, err := SetupMockDatabase()
    if err != nil {
        t.Fatalf("Error setting up mock database: %v", err)
    }
    defer TeardownMockDatabase(mockDB) // Cleanup after test execution

    // Insert test data into the mock database
    if err := mockDB.InsertTestData(); err != nil {
        t.Fatalf("Error inserting test data: %v", err)
    }

    // Run tests...
}


// pkg/api/handlers/organization_handler.go

package handlers

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/yourusername/yourproject/pkg/shared"
)

type OrganizationRepository struct {
    collection *mongo.Collection
}

func NewOrganizationRepository(collection *mongo.Collection) *OrganizationRepository {
    return &OrganizationRepository{collection}
}

// Implement CRUD operations for organizations

// Use the shared logging function
func SomeFunction() {
    // Example error to log
    someError := "Example error occurred"
    shared.LogError(someError)
}


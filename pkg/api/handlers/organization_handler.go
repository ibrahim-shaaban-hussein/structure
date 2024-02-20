// pkg/api/handlers/organization_handler.go

package handlers

import (
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/ibrahim-shaaban-hussein/structure/pkg/shared"
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


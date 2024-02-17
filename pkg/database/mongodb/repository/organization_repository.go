package repository

import (
    "context"

    "go.mongodb.org/mongo-driver/mongo"
)

type OrganizationRepository struct {
    collection *mongo.Collection
}

func NewOrganizationRepository(collection *mongo.Collection) *OrganizationRepository {
    return &OrganizationRepository{collection}
}

// Implement CRUD operations for organizations


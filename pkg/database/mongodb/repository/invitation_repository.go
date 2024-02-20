package repository

import (
    "go.mongodb.org/mongo-driver/mongo"
)

type InvitationRepository struct {
    collection *mongo.Collection
}

func NewInvitationRepository(collection *mongo.Collection) *InvitationRepository {
    return &InvitationRepository{collection}
}

// Implement CRUD operations for invitations


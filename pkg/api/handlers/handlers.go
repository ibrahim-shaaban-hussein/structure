package handlers

import (
    "fmt"
)

// UserService defines the interface for user-related operations.
type UserService interface {
    GetUser(userID string) (User, error)
}

// OrganizationService defines the interface for organization-related operations.
type OrganizationService interface {
    UpdateOrganization(orgID string, name string) error
}

// User represents user data.
type User struct {
    ID   string
    Name string
    // Add more fields as needed
}

// OrganizationHandler depends on UserService and OrganizationService interfaces.
type OrganizationHandler struct {
    UserService        UserService
    OrganizationService OrganizationService
}

// GetUserHandler retrieves user data using UserService.
func (h *OrganizationHandler) GetUserHandler(userID string) (User, error) {
    return h.UserService.GetUser(userID)
}

// UpdateOrganizationHandler updates organization data using OrganizationService.
func (h *OrganizationHandler) UpdateOrganizationHandler(orgID string, name string) error {
    err := h.OrganizationService.UpdateOrganization(orgID, name)
    if err != nil {
        return fmt.Errorf("failed to update organization: %v", err)
    }
    return nil
}


package handlers

// UserHandler is responsible for handling user-related operations.
type UserHandler struct {
    // UserService interface is not redeclared here
    UserService UserService
}

// GetUser retrieves user data using UserService.
func (h *UserHandler) GetUser(userID string) (User, error) {
    return h.UserService.GetUser(userID)
}


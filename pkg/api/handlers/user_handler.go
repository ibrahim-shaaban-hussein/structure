// user_handler.go

package handlers

// UserService defines the interface for user-related operations.
type UserService interface {
    GetUser(userID string) (User, error)
}

// UserHandler depends on UserService interface.
type UserHandler struct {
    UserService UserService
}

// GetUser retrieves user data using UserService.
func (h *UserHandler) GetUser(userID string) (User, error) {
    return h.UserService.GetUser(userID)
}


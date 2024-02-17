package handlers

// UserService defines the interface for user-related operations.
type UserService interface {
	GetUser(userID string) (User, error)
}

// User represents user data.
type User struct {
	ID   string
	Name string
	// Add more fields as needed
}

// UserHandler depends on UserService interface.
type UserHandler struct {
	UserService UserService
}

// GetUserHandler retrieves user data using UserService.
func (h *UserHandler) GetUserHandler(userID string) (User, error) {
	return h.UserService.GetUser(userID)
}


package middleware

// isValidToken is a function to validate the token
func isValidToken(token string) bool {
    // Example token validation logic (replace with your actual implementation)
    // For demonstration purposes, we'll assume the token is valid if its length is greater than 10 characters
    import "github.com/ibrahims.hussein/structure/pkg/api/handlers/middleware/auth"
    return len(token) > 10
}


package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Define a struct representing the expected JSON format of the request body for signup
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func SigninHandler(c *gin.Context) {
	var req SigninRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// As example: we check if the user exists and the password matches
	// I will replace this with the real actual authentication logic
	if req.Email != "user@example.com" || req.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT tokens
	accessToken, err := generateAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Return success response with tokens
	resp := SigninResponse{
		Message:      "Signin successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, resp)
}

// Function to generate access token
func generateAccessToken() (string, error) {
	// Create a new JWT token for the access token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires in 1 hour

	// Sign token with secret key and return
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Function to generate refresh token
func generateRefreshToken() (string, error) {
	// Similar to generateAccessToken but with longer expiration time
	// I might also store the refresh token in the database for revocation if needed
	// Return the token string
	return "refresh-token", nil
}

// Handler function for the /signup endpoint
func SignupHandler(c *gin.Context) {
	// Initialize an empty struct to hold the parsed JSON data
	var req SignupRequest

	// Parse the request body and bind it to the struct
	if err := c.BindJSON(&req); err != nil {
		// If there's an error parsing the JSON, return a bad request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Access the fields of the struct to retrieve the required information
	name := req.Name
	email := req.Email
	password := req.Password

	// Perform further processing with the extracted information
	// For example, you can validate the input, create a user, etc.

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully", "name": name, "email": email})
}

func main() {
	r := gin.Default()

	// Define route for signin
	r.POST("/signin", SigninHandler)

	// Define route for signup
	r.POST("/signup", SignupHandler)

	// Run the server
	r.Run(":8080")
}

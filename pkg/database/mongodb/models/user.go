package models

type User struct {
    ID       string `json:"id" bson:"_id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    // Add other fields as needed
}


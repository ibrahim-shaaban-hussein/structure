package models

type Invitation struct {
    ID           string `json:"id" bson:"_id"`
    Organization string `json:"organization"`
    InvitedUser  string `json:"invited_user"`
    // Add other fields as needed
}


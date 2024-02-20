package models

type Organization struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name"`
	// Add other fields as needed
}

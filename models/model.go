package models


import "time"


type Model struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	DateOfBirth string    `json:"date_of_birth" bson:"date_of_birth"`
	Address     string    `json:"address" bson:"address"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
} 


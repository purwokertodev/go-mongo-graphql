package model

import (
	"time"
)

//Profile
type Profile struct {
	ID        string    `json:"id" bson:"id"`
	FirstName string    `json:"firstName" bson:"first_name"`
	LastName  string    `json:"lastName" bson:"last_name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

//Profiles
type Profiles []Profile

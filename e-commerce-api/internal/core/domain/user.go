package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	FullName     string    `json:"full_name"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"user_password"`
}

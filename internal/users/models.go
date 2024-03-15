package users

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	Active       bool
}
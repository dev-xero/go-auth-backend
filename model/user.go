package model

import "github.com/google/uuid"

/*
User model struct

Fields:
  - ID:       uuid
  - Username: string
  - Email:    string
  - Password: string
*/
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

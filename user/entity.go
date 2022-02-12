package user

import "time"

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	PhoneNumber  string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

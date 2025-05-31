package entities

import (
	"gorm.io/gorm"
)

// User represents a user in the authentication system.
type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
}

package repository

import (
	"context"

	"github.com/himakhaitan/noreboothq/services/auth/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}

// userRepository implements UserRepository interface for user-related database operations.
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// GetByEmail retrieves a user by their email address from the database.
// It returns the user if found, or an error if not found or if there is a database error.
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

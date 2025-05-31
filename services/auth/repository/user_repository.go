package repository

import (
	"context"

	"github.com/himakhaitan/noreboothq/services/auth/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

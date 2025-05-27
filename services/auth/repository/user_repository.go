package repository

import (
	"context"

	"github.com/himakhaitan/noreboothq/services/auth/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
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

func (r *userRepository) UpdateLastLogin(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Model(&entities.User{}).
		Where("id = ?", userID).
		Update("last_login_at", gorm.Expr("NOW()")).Error
}

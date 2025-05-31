package controllers

import "github.com/himakhaitan/noreboothq/services/auth/repository"

// AuthController handles authentication-related operations.
// It interacts with the UserRepository to perform user-related actions such as login, etc.
type AuthController struct {
	userRepo *repository.UserRepository
}

// NewAuthController creates a new instance of AuthController with the provided UserRepository.
func NewAuthController(userRepo *repository.UserRepository) *AuthController {
	return &AuthController{userRepo: userRepo}
}

// func (c *AuthController) Login(ctx context.Context, email string, password string) (string, error) {
//  // Handle user login logic here
// }

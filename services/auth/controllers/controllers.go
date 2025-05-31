package controllers

import "github.com/himakhaitan/noreboothq/services/auth/repository"

type AuthController struct {
	userRepo *repository.UserRepository
}

func NewAuthController(userRepo *repository.UserRepository) *AuthController {
	return &AuthController{userRepo: userRepo}
}

// func (c *AuthController) Login(ctx context.Context, email string, password string) (string, error) {
//  // Handle user login logic here
// }

package handlers

import (
	authpb "github.com/himakhaitan/noreboothq/proto/auth"
	"github.com/himakhaitan/noreboothq/services/auth/controllers"
	"go.uber.org/zap"
)

// Package handlers provides the HTTP handlers for the authentication service.
type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	ctrl   *controllers.AuthController
	logger *zap.Logger
}

// NewAuthHandler creates a new instance of AuthHandler with the provided AuthController and logger.
func NewAuthHandler(ctrl *controllers.AuthController, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		ctrl:   ctrl,
		logger: logger,
	}
}

// func (h *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
// 	h.logger.Info("Login request received", zap.String("email", req.Email))
// 	token, err := h.ctrl.Login(ctx, req.Email, req.Password)

// }

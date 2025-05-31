package handlers

import (
	authpb "github.com/himakhaitan/noreboothq/proto/auth"
	"github.com/himakhaitan/noreboothq/services/auth/controllers"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	ctrl   *controllers.AuthController
	logger *zap.Logger
}

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

package handlers

import (
	authpb "github.com/himakhaitan/noreboothq/proto/auth"
	"go.uber.org/zap"
	"github.com/himakhaitan/noreboothq/services/auth/controllers"
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

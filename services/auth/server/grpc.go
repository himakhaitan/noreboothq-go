package server

import (
	"fmt"
	"net"

	authpb "github.com/himakhaitan/noreboothq/proto/auth"
	"github.com/himakhaitan/noreboothq/services/auth/controllers"
	"github.com/himakhaitan/noreboothq/services/auth/handlers"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	grpcServer *grpc.Server
	logger     *zap.Logger
	port       int
}

func NewGRPCServer(logger *zap.Logger, port int) *GRPCServer {
	return &GRPCServer{
		grpcServer: grpc.NewServer(),
		logger:     logger,
		port:       port,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	// Create controller and handler
	authCtrl := controllers.NewAuthController()
	authHandler := handlers.NewAuthHandler(authCtrl, s.logger)

	// Register gRPC service
	authpb.RegisterAuthServiceServer(s.grpcServer, authHandler)

	s.logger.Info("gRPC server started", zap.Int("port", s.port))
	return s.grpcServer.Serve(lis)
}

package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	authpb "github.com/himakhaitan/noreboothq/proto/auth"
	"github.com/himakhaitan/noreboothq/services/auth/controllers"
	"github.com/himakhaitan/noreboothq/services/auth/handlers"
	"github.com/himakhaitan/noreboothq/services/auth/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	grpcServer *grpc.Server
	logger     *zap.Logger
	port       int
	userRepo   *repository.UserRepository
}

func NewGRPCServer(logger *zap.Logger, userRepo *repository.UserRepository, port int) *GRPCServer {
	return &GRPCServer{
		grpcServer: grpc.NewServer(),
		logger:     logger,
		port:       port,
		userRepo:   userRepo,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	authCtrl := controllers.NewAuthController(s.userRepo)
	authHandler := handlers.NewAuthHandler(authCtrl, s.logger)

	authpb.RegisterAuthServiceServer(s.grpcServer, authHandler)

	// Channel to listen for errors from grpcServer.Serve
	serveErrCh := make(chan error, 1)

	go func() {
		s.logger.Info("gRPC server started", zap.Int("port", s.port))
		serveErrCh <- s.grpcServer.Serve(lis)
	}()

	// Channel to listen for OS signals
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-stopCh:
		s.logger.Info("Shutdown signal received", zap.String("signal", sig.String()))

		// Graceful stop with timeout context
		doneCh := make(chan struct{})
		go func() {
			s.grpcServer.GracefulStop()
			close(doneCh)
		}()

		// Wait for graceful stop or timeout
		select {
		case <-doneCh:
			s.logger.Info("gRPC server stopped gracefully")
		case <-time.After(10 * time.Second):
			s.logger.Warn("Timeout reached, forcing gRPC server stop")
			s.grpcServer.Stop()
		}
		return nil

	case err := <-serveErrCh:
		return err
	}
}

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/himakhaitan/noreboothq/services/auth/config"
	"github.com/himakhaitan/noreboothq/services/auth/entities"
	"github.com/himakhaitan/noreboothq/services/auth/repository"
	"github.com/himakhaitan/noreboothq/services/auth/server"
	sharedConfig "github.com/himakhaitan/noreboothq/shared/config"
	sharedDB "github.com/himakhaitan/noreboothq/shared/db"
	"github.com/himakhaitan/noreboothq/shared/env"
	sharedLogger "github.com/himakhaitan/noreboothq/shared/logger"
	"go.uber.org/zap"
)

func main() {
	// Default values to use if none provided
	defaults := struct {
		configPath string
		env        string
	}{
		configPath: "services/auth/config",
		env:        "development",
	}
	resolved := env.ResolveEnvConfig(defaults.configPath, defaults.env)
	cfg, err := sharedConfig.LoadConfig[config.AuthServiceConfig](resolved.ConfigPath, resolved.Env)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	// Initialize the logger with the loaded configuration
	err = sharedLogger.Init(
		sharedLogger.Config{
			ServiceName: "auth-service",
			Environment: resolved.Env,
		}, cfg.Log.Level,
	)
	if err != nil {
		panic("failed to init logger: " + err.Error())
	}
	defer sharedLogger.Sync() // flushes logs on exit

	// Initialize database connection with User model for migration
	// You can add more models as needed
	db, err := sharedDB.NewConnection(sharedDB.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	}, sharedLogger.Logger(), &entities.User{})
	if err != nil {
		sharedLogger.Logger().Fatal("Failed to connect to database", zap.Error(err))
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	sharedLogger.Logger().Info("Auth Service Started")

	// Graceful shutdown context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start the gRPC server
	grpcServer := server.NewGRPCServer(sharedLogger.Logger(), &userRepo, cfg.Server.Port)
	if err := grpcServer.Start(ctx); err != nil {
		sharedLogger.Logger().Fatal("Failed to start gRPC server", zap.Error(err))
	}
}

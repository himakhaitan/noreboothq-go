package main

import (
	"github.com/himakhaitan/noreboothq/services/auth/config"
	"github.com/himakhaitan/noreboothq/services/auth/server"
	sharedConfig "github.com/himakhaitan/noreboothq/shared/config"
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

	sharedLogger.Logger().Info("Auth Service Started")

	grpcServer := server.NewGRPCServer(sharedLogger.Logger(), cfg.Server.Port)
	if err := grpcServer.Start(); err != nil {
		sharedLogger.Logger().Fatal("Failed to start gRPC server", zap.Error(err))
	}
}

package db

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection creates a new database connection and runs migrations if models are provided
func NewConnection(config Config, zapLogger *zap.Logger, models ...interface{}) (*gorm.DB, error) {
	// Set default SSLMode if not provided
	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	// Create custom GORM logger using Zap
	gormLogger := NewZapLogger(zapLogger)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations if models are provided
	if len(models) > 0 {
		if err := db.AutoMigrate(models...); err != nil {
			return nil, fmt.Errorf("failed to run migrations: %w", err)
		}
	}

	return db, nil
}

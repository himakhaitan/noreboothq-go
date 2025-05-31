package db

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// ZapLogger is a custom GORM logger that uses Zap
type ZapLogger struct {
	ZapLogger *zap.Logger
	Config    gormlogger.Config
}

// NewZapLogger creates a new ZapLogger instance
func NewZapLogger(zapLogger *zap.Logger) *ZapLogger {
	return &ZapLogger{
		ZapLogger: zapLogger,
		Config: gormlogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  gormlogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	}
}

// LogMode implements the gormlogger.Interface
func (l *ZapLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.Config.LogLevel = level
	return &newLogger
}

// Info implements the gormlogger.Interface
func (l *ZapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.Config.LogLevel >= gormlogger.Info {
		l.ZapLogger.Sugar().Infof(msg, data...)
	}
}

// Warn implements the gormlogger.Interface
func (l *ZapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.Config.LogLevel >= gormlogger.Warn {
		l.ZapLogger.Sugar().Warnf(msg, data...)
	}
}

// Error implements the gormlogger.Interface
func (l *ZapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.Config.LogLevel >= gormlogger.Error {
		l.ZapLogger.Sugar().Errorf(msg, data...)
	}
}

// Trace implements the gormlogger.Interface
func (l *ZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Config.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.Duration("elapsed", elapsed),
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.Config.IgnoreRecordNotFoundError) {
		fields = append(fields, zap.Error(err))
		l.ZapLogger.Error("gorm trace", fields...)
		return
	}

	if l.Config.SlowThreshold != 0 && elapsed > l.Config.SlowThreshold {
		l.ZapLogger.Warn("gorm slow query", fields...)
		return
	}

	l.ZapLogger.Debug("gorm trace", fields...)
}

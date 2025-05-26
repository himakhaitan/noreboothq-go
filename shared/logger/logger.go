package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	ServiceName string
	Environment string
	Port        int
}

// package-level logger instance (private)
var logger *zap.Logger

// Init initializes the package-level logger.
func Init(cfg Config, logLevel string) error {
	isProd := strings.EqualFold(cfg.Environment, "production")

	level := zapcore.InfoLevel
	if logLevel != "" {
		if err := level.UnmarshalText([]byte(strings.ToLower(logLevel))); err != nil {
			level = zapcore.InfoLevel
		}
	} else {
		if !isProd {
			level = zapcore.DebugLevel
		}
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if isProd {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		level,
	)

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	logger = zap.New(core, options...).With(
		zap.String("service", cfg.ServiceName),
		zap.String("env", cfg.Environment),
		zap.Int("port", cfg.Port),
	)

	return nil
}

// Logger returns the package-level logger instance.
// Panics if Init is not called.
func Logger() *zap.Logger {
	if logger == nil {
		panic("logger: Init must be called before using Logger()")
	}
	return logger
}

// Sync flushes buffered logs.
func Sync() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}

# `shared/logger`

## 📦 Overview

This package provides a reusable logging utility built on top of `zap` — Uber's high-performance structured logging library.

It enables services to initialize a consistent, configurable logger with sensible defaults for both local development and production environments.

## 🧩 Folder Structure

```bash
shared/logger/
└── logger.go  # Zap-based logger with environment-aware output
```

## 🛠️ What It Does

- Provides a package-level logger for consistent access across your service.
- Supports both JSON (for production) and console (for local/dev) formats.
- Allows per-service labeling (service, env) for log observability.
- Sets log levels dynamically (e.g., debug, info, error).
- Handles panic if logger is accessed before initialization.

## ⚙️ How to Use

1. **Initialize the Logger**
Call `logger.Init(...)` early in your service startup, typically in `main.go`:

```go
import "github.com/himakhaitan/noreboothq/shared/logger"

err := logger.Init(logger.Config{
    ServiceName: "auth-service",
    Environment: "dev", // or "production"
}, "debug")

if err != nil {
    panic("failed to initialize logger: " + err.Error())
}
```

2. **Get the Logger**

```go
log := logger.Logger()
log.Info("starting service...")
```

3. **Flush Before Exit**

```go
defer logger.Sync()
```

## 🔍 Output Behavior 

| Mode       | Output Format | Log Level |
| ---------- | ------------- | --------- |
| dev        | Console       | debug     |
| production | JSON          | info      |

## 💡 Example Log (Development)

```css
2025-05-31T16:48:19.143+0530    info    main.go:14  starting service...      {"service": "auth-service", "env": "dev"}
```

## 🧼 Best Practices

- Always call Init() before using Logger(). It panics if uninitialized.
- Use structured logging: log.Info("msg", zap.String("key", "value"))
- Set log level via config or env for flexibility.
- Use logger.Sync() to flush logs on shutdown, especially in production.
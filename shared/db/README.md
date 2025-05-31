# `shared/db`

## 📦 Overview

This package provides a shared utility to manage PostgreSQL database connections across services using [GORM](https://gorm.io/) as the ORM and [Zap](https://pkg.go.dev/go.uber.org/zap) as the structured logger.

It abstracts away connection boilerplate, integrates automatic migrations, and attaches a custom GORM logger for consistent logging across your application.

## 🧩 Folder Structure

```vbnet
shared/db/
├── connection.go   # Initializes GORM connection with optional migrations
└── logger.go       # Custom GORM logger using Zap
```

## 🔌 Features

- 🔄 Centralized database connection setup
- 🧪 Optional model-based migrations
- 📊 Structured GORM query logging via Zap
- ⚠️ Slow query detection
- 🧼 Ignores common noise like record not found errors

## 🛠️ Usage

1. **Define a Config Struct (usually in `config/types.go`)**

```go
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}
```

2. **Initialize the Connection**
In your `main.go` or `server.go`:

```go
import (
    "github.com/himakhaitan/noreboothq/shared/db"
    "go.uber.org/zap"
)

cfg := db.Config{
    Host:     "localhost",
    Port:     5432,
    User:     "postgres",
    Password: "password",
    DBName:   "noreboothq",
    SSLMode:  "disable",
}

zapLogger, _ := zap.NewProduction()

dbConn, err := db.NewConnection(cfg, zapLogger)
if err != nil {
    log.Fatalf("failed to connect: %v", err)
}
```

3. **Optional: Auto-Migrate Models**

You can also pass models directly into `NewConnection()` to run `AutoMigrate`:

```go
dbConn, err := db.NewConnection(cfg, zapLogger, &User{}, &Session{})
```

This is useful for initial schema setup or test environments.

## 🧩 GORM + Zap Logger

The custom logger logs:
- Query execution time
- SQL query string
- Rows affected
- Errors (if any)
- Warnings for slow queries

Example log (in JSON if in prod):

```json
{
  "level": "debug",
  "msg": "gorm trace",
  "sql": "SELECT * FROM users WHERE id = 1",
  "rows": 1,
  "elapsed": "3.2ms"
}
```

## 🔍 Behavior Details

| Feature                     | Description                        |
| --------------------------- | ---------------------------------- |
| `SlowThreshold`             | 200ms; logs slow queries as `WARN` |
| `IgnoreRecordNotFoundError` | true; avoids spamming logs         |
| `LogLevel`                  | Set to `Info` by default           |
| `Colorful`                  | false (terminal-agnostic output)   |


## ✅ When to Use

- Shared DB setup across microservices
- Consistent GORM logging
- Simplified migration support

## 🧠 Good to Know

- Supports PostgreSQL only (but can be extended for others).
- Designed for internal service use; not a general-purpose DB abstraction.
- Zap logger must be initialized before calling NewConnection.
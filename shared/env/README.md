# `shared/env`

## 📦 Overview

This package provides a standardized utility to resolve environment-specific runtime configuration. It simplifies the process of managing runtime env and config path values using a clear precedence order:

```objectivec
CLI flags > Environment variables > Defaults
```

This ensures your services can be easily run in different environments (dev, staging, prod) without hardcoding values or repeating setup logic.

## 🧩 Folder Structure

```bash
shared/env/
└── env.go  # Resolves runtime environment and config path
```

## 🛠️ What It Does

- Parses command-line flags (--env, --config) if provided.
- Falls back to environment variables (ENV, CONFIG_PATH) if flags aren't set.
- Uses hardcoded defaults if neither flags nor env vars are present.
- Returns a struct containing both resolved values.

## ⚙️ How to Use

1. **Import and Call**

Typically called at the top of your service's `main.go`:

```go
import "github.com/himakhaitan/noreboothq/shared/env"

cfg := env.ResolveEnvConfig("config", "development")
fmt.Println(cfg.Env)         // e.g., "production"
fmt.Println(cfg.ConfigPath)  // e.g., "./config"
```

2. **CLI Flags**

You can override values using flags:

```bash
go run main.go --env=production --config=/etc/my-service
```

3. **Environment Variables**

If flags aren't set, you can use env vars:

```bash
export ENV=staging
export CONFIG_PATH=./configs/staging
go run main.go
```

## ✅ Precedence Order

| Source            | Priority |
| ----------------- | -------- |
| CLI Flags         | Highest  |
| Env Vars          | Medium   |
| Default Fallbacks | Lowest   |

## 📄 Return Type

```go
type EnvConfig struct {
    Env        string // Resolved environment
    ConfigPath string // Path to YAML config files
}
```

## 💡 Best Practices

- Use this early in service startup to configure logger, load configs, etc.
- Keep default values minimal; prefer environment variables or CLI flags for flexibility.
- Make sure to call flag.Parse() only once in your program.
# ⚙️ `config/` — Configuration for Auth Service

This folder contains all the configuration definitions and files used by the Auth Service.

## 📁 Contents

- `types.go` - Go types used to unmarshal config values loaded at runtime
- `base.yaml` — Base configuration shared across environments
- `production.yaml` — Environment-specific overrides for production
- `staging.yaml` — Environment-specific overrides for staging

## 🔄 How It Works

Configuration is loaded using the `shared/config` loader. It:

1. Loads the common `base.yaml`
2. Then overlays it with the selected environment file (e.g., `production.yaml`)
3. Populates the `AuthServiceConfig` Go struct

## 🧪 Example Usage

```go
cfg, err := sharedConfig.LoadConfig[config.AuthServiceConfig](path, env)
```

## 🏗 Structure

Here's how the configuration types are structured:

```go
type AuthServiceConfig struct {
  Server ServerConfig
  JWT    JWTConfig
  Log    LogConfig
  DB     DatabaseConfig
}
```

This allows centralized and environment-specific configuration management for the Auth Service.
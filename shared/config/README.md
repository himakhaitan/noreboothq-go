# `shared/config`

## 📦 Overview

This package provides a generic configuration loader using the `koanf`library.
It allows services to load and merge YAML-based configurations — both base and environment-specific — into strongly typed Go structs.

The goal is to provide a central, reusable config loader for all services without duplicating logic.

## 🛠️ What It Does

Loads a base config file: `base.yaml`

Optionally merges an environment-specific override: `<env>.yaml`

Unmarshals the final result into a strongly typed struct defined in the calling service

## 🧩 Folder Structure

```arduino
shared/config/
└── loader.go     # Core config loading logic using koanf
```

## 🔍 How It Works

The core function is:

```go
func LoadConfig[T any](basePath string, env string) (*T, error)
```

- `T` is the config type defined in your service (e.g. `AuthServiceConfig`)
- `basePath` is the directory containing your YAML files
- `env` is the environment name (e.g., `"dev"`, `"prod"`), used to load `dev.yaml`, `prod.yaml`, etc.

## 🧪 Example Usage (from `auth-service`)

1. Define config type in `services/auth/config/types.go`:

```go
type AuthServiceConfig struct {
	Server ServerConfig   `koanf:"server"`
	JWT    JWTConfig      `koanf:"jwt"`
	Log    LogConfig      `koanf:"logging"`
	DB     DatabaseConfig `koanf:"database"`
}
```

2. Load the config:

```go
cfg, err := config.LoadConfig[AuthServiceConfig]("configs/auth", "dev")
if err != nil {
  log.Fatalf("Failed to load config: %v", err)
}
```

This would merge:

```bash
configs/auth/base.yaml
configs/auth/dev.yaml
```

And return a fully populated `AuthServiceConfig`.

## 🧼 Best Practices

- Define your config schema in a service-specific types.go file.
- Use koanf tags (koanf:"key") to bind YAML keys to struct fields.
- Keep base YAML files for defaults; override per environment.
- Do not hardcode config values — use YAML + env layering.

## 📁 YAML Layout Sample

```yaml
# base.yaml
server:
  port: 8080
jwt:
  secret_key: "supersecret"

# dev.yaml
server:
  port: 8081
```
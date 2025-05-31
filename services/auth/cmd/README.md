# `cmd/` — Service Entrypoint

This folder contains the main entrypoint for the Auth Service.

## 📌 Purpose

The `main.go` file is responsible for:

- Loading environment-specific configuration files
- Initializing structured logging
- Establishing the database connection and running migrations
- Setting up repositories and other core dependencies
- Starting the gRPC server

## 🧪 How to Run

```bash
go run services/auth/cmd/main.go \
  --env=development \
  --config=services/auth/config
```

Alternatively, you can set env variables instead of flags:

```bash
export ENV=development
export CONFIG_PATH=services/auth/config

go run services/auth/cmd/main.go
```

## ⚙️ Dependencies Used

- 🧾 `shared/config` – Loads and merges base + env config files
- 🌍 `shared/env` – Resolves config/env values from flags/env vars
- 🪵 `shared/logger` – Structured logging with Zap
- 🛢 `shared/db` – GORM DB connection and migration runner
- 🔒 `services/auth/repository` – User repository
- 🎯 `services/auth/server` – gRPC server and service wiring
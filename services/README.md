# `services`

This folder contains all the microservices that make up the full NoRebootHQ platform. Each service is organized as a clean, standalone unit with its own entrypoint, configuration, server wiring, handlers, repositories, and business logic.

## 🔍 Purpose

Each subfolder inside `services/` represents a single microservice — such as:

- `auth`: handles authentication logic
- `identity`: manages users, orgs, and roles
- `config`: deals with configuration versioning and environment-specific values
- `deployment`: delivers live configuration updates to consumers

These services do not share logic directly — instead, they communicate via gRPC and consume shared utilities from the `shared/` folder.

## 🧱 Common Folder Structure

Every service follows a consistent internal layout:

```bash
services/
└── [service-name]/
    ├── cmd/           # Entrypoint (main.go)
    ├── config/        # Service-specific configuration types and files
    ├── controllers/   # Business logic layer
    ├── entities/      # Domain models & DB schema definitions
    ├── handlers/      # gRPC/HTTP request handlers
    ├── repository/    # Data access layer (DB interactions)
    └── server/        # gRPC server setup and wiring
```

## 🧰 Reusability

Each service makes use of the shared utilities (from `/shared`) for:

- Config loading
- Env resolution
- Database connection
- Logging setup

This promotes consistency across services while preserving full separation of concerns.

## 🚀 Usage

These services are designed to be run independently but can be composed together in a production setup using container orchestration (e.g. Docker Compose, Kubernetes, etc.).

Each one has its own `main.go` that can be launched individually and is testable in isolation.
# `shared`

## 📦 Overview

The `shared/` folder **contains common utilities and abstractions** that are reused across multiple services in the project.

These are building blocks that help keep code DRY (Don't Repeat Yourself), standardized, and easier to maintain across the codebase.

Typical contents of this folder include:
- Configuration loaders
- Environment variable resolution helpers
- Database connection setup
- Logger initialization

Each of these modules is designed to be importable and used directly by any service, reducing duplication and enforcing consistency in implementation.

## 🧩 Structure

```bash
shared/
├── config/        # Load and parse YAML configs
├── env/           # Load and resolve environment-specific values
├── db/            # DB connection setup and lifecycle handling
└── logger/        # Centralized logger configuration
```

> Note: Each package inside `shared/` have its own README explaining specifics and usage examples.

## ✅ When to Use

Use the `shared/` modules when:
- You need a common setup across services (e.g., logger, DB, configs)
- You're bootstrapping a new service and want consistent defaults
- You're writing code that would otherwise be copied in multiple places

## 🔄 Import Example

In a Go service:
```go
import (
  "github.com/himakhaitan/noreboothq/shared/config"
  "github.com/himakhaitan/noreboothq/shared/logger"
)
```

## 🧼 Best Practices

- Keep only generic, reusable logic here — service-specific code should live inside the service itself.
- Avoid tight coupling to a specific service’s context or schema.
- Maintain clean, minimal APIs to make reuse easy and intuitive.
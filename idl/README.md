# `idl`

## 📋 Overview

The `idl` folder (short for Interface Definition Language) contains the `.proto` files that define all the **gRPC service contracts**, message structures, and API interfaces used across the project.

These definitions are the source of truth for how services communicate — both internally (between microservices) and externally (e.g., SDK clients, gateways).

By storing them centrally, we ensure consistent, version-controlled communication interfaces across all services.

## 🔧 What Are `.proto` Files?

`.proto` files are written using **Protocol Buffers (Proto3)** — a language-neutral, platform-neutral interface definition language developed by Google.

These files:
- Define **services** and **RPC methods**
- Define the **request/response message structures**
- Are used to generate strongly-typed client and server code in Go (and other languages)

This ensures type-safe, schema-driven communication between systems.

## 🧩 Folder Structure

```csharp
idl/
├── auth/
│   └── auth.proto   ← Defines the AuthService interface
```

Each subfolder under `idl/` represents a domain or microservice boundary (e.g., `auth`, `config`, `user`, etc.).

This keeps `.proto` files modular and aligned with service responsibilities.

## 📝 Sample: `auth.proto`

```proto
syntax = "proto3";

package auth;

option go_package = "github.com/himakhaitan/noreboothq/proto/auth;authpb";

service AuthService {
    rpc Login(LoginRequest) returns (LoginResponse);
}

message LoginRequest {
  string email = 1;
  string password = 2;
}

message LoginResponse {
  string access_token = 1;
  string token_type = 2; // e.g., "Bearer"
  int64 expires_in = 3; // in seconds
}
```

## 🔍 Key Components

| Line/Block                               | Purpose                                                                      |
| ---------------------------------------- | ---------------------------------------------------------------------------- |
| `syntax = "proto3";`                     | Specifies we're using Proto3 syntax (the latest and recommended version)     |
| `package auth;`                          | Defines the protobuf package namespace. Helps organize messages and services |
| `option go_package = ...;`               | Specifies the Go import path for generated code, and alias (`authpb`)        |
| `service AuthService`                    | Defines a gRPC service called `AuthService`                                  |
| `rpc Login(...)`                         | Defines a single RPC method `Login` with its request and response messages   |
| `message LoginRequest` / `LoginResponse` | Define the schema for input and output of the `Login` RPC                    |

## ⚙️ Code Generation

Once `.proto` files are created, you need to generate Go code using:

```bash
./scripts/proto.sh
```

This will output:
```go
proto/
└── auth/
    ├── auth.pb.go
    └── auth_grpc.pb.go
```

These generated files contain:
- Structs for request/response messages
- Server and client interfaces for gRPC services

## 📦 Best Practices

- Keep `.proto` files minimal and domain-specific (one service per file)
- Use clear, versioned Go import paths in `go_package` (as shown above)
- Document message fields with comments where helpful
- Always regenerate code after modifying `.proto` files
- Avoid breaking changes unless you’re versioning services explicitly

## 🔗 Related

- `proto`/ — Output folder where generated `.pb.go` files are stored
- `scripts/proto.sh` — Automation script that generates Go code from `.proto` files
- `services/` — Services that import and use the generated gRPC interfaces

## 📚 References

- [Protocol Buffers Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)
- [gRPC in Go](https://grpc.io/docs/languages/go/quickstart/)
- [Google Protobuf Style Guide](https://developers.google.com/protocol-buffers/docs/style)


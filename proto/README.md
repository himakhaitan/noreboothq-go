# `proto`

## 📋 Overview

The `proto/` folder contains **auto-generated Go code** from the `.proto` files located in the `idl/` directory.

These files are produced by the `proto.sh` script and are required to enable gRPC communication between services using type-safe, strongly-typed interfaces.

You should **not manually edit** anything in this folder — always regenerate them via the script whenever changes are made to your `.proto` definitions.

## 🧬 Folder Structure

```bash
proto/
└── auth/
    ├── auth.pb.go         # Message types, helpers, and marshaling logic
    └── auth_grpc.pb.go    # gRPC server & client interfaces
```

Each subfolder here mirrors the package structure defined in your `.proto` files (see `option go_package`).

For example, the `auth.proto` file defines:

```proto
option go_package = "github.com/himakhaitan/noreboothq/proto/auth;authpb";
```

Which means the generated Go code will be output under `proto/auth` and imported as `authpb`.

## ⚙️ How This Code Gets Here

Run the following command to generate code:

```bash
./scripts/proto.sh
```

This script:
- Compiles all .proto files in idl/
- Generates Go code into the proto/ directory
- Ensures the structure aligns with Go import paths for seamless usage in services

## 🧠 What's Inside the Generated Files

| File           | Description                                                                                              |
| -------------- | -------------------------------------------------------------------------------------------------------- |
| `*.pb.go`      | Contains Go structs for all messages (e.g., `LoginRequest`, `LoginResponse`)                             |
| `*_grpc.pb.go` | Contains gRPC server and client interfaces for services (e.g., `AuthServiceServer`, `AuthServiceClient`) |

These files are used directly in the services to:
- Register gRPC servers
- Implement RPC handlers
- Make inter-service calls using generated clients

## 📦 Usage Example

In a service like `auth-service`, you would import the generated client like this:

```go
import authpb "github.com/himakhaitan/noreboothq/proto/auth"
```

And register your service like

```go
authpb.RegisterAuthServiceServer(grpcServer, myAuthHandler)
```

## 📌 Good to Know

- Always regenerate the proto code after changing .proto files.
- Never manually modify the contents of this folder.
- Use versioned or consistent import paths in option go_package to avoid import issues.

## 🛠️ Troubleshooting

If you see issues like “missing proto imports” or “undeclared services,” re-run:

```bash
./scripts/proto.sh
```

And make sure your Go environment includes the necessary `protoc-gen-go` and `protoc-gen-go-grpc` plugins.

## 🔗 Related

- `idl/` — Contains the actual .proto definitions.
- `scripts/proto.sh` — The script that compiles .proto files into this folder.
- `services/` — Services that use the generated code to implement business logic.
# 🧾 `server/` — gRPC Server Layer

This layer is responsible for setting up, starting, and managing the lifecycle of the gRPC server that exposes the Auth service.

## 📁 Contents

`grpc.go` — Defines the `GRPCServer` struct that encapsulates the gRPC server, its configuration, and lifecycle methods.

## 🧠 Purpose

The server layer handles:
- 🔌 Creating and configuring the underlying gRPC server
- 🎛️ Registering service handlers (e.g., AuthHandler)
- 📡 Listening on the configured TCP port for incoming requests
- 🛑 Managing graceful shutdown on context cancellation or system signals
- 📋 Logging server events like start, stop, errors, and shutdown progress

## 🧱 Example

```go
grpcServer := server.NewGRPCServer(logger, userRepo, port)
if err := grpcServer.Start(ctx); err != nil {
    logger.Fatal("Failed to start gRPC server", zap.Error(err))
}
```

The `Start` method blocks and runs the server until the context is canceled or an error occurs, ensuring clean shutdown and resource cleanup.
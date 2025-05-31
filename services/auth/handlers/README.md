# 🧾 `handlers/` — gRPC Request Handlers

This layer defines the gRPC handlers that connect external requests to the internal business logic provided by controllers.

## 📁 Contents

- `handlers.go` — Contains the `AuthHandler` which implements the `AuthService` gRPC server defined in the protobuf definition.

## 🧠 Purpose

Handlers are the entry point for gRPC requests. Their responsibilities include:
- 🔌 Receiving incoming gRPC requests
- 📋 Validating and logging request metadata
- 🎛️ Delegating to the appropriate controller
- 📤 Returning gRPC responses

## 🧱 Example

```go
type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	ctrl   *controllers.AuthController
	logger *zap.Logger
}
```

You can register this handler to your gRPC server like so:

```go
authpb.RegisterAuthServiceServer(grpcServer, handler)
```

## 🛠️ In Progress

Example method stubs like `Login(ctx, req)` can be implemented by:

1. Logging the request
2. Passing data to `ctrl.Login(...)`
3. Returning a valid gRPC `LoginResponse`
# 🧬 `entities/` — Data Models

This folder contains the database models used in the Auth Service. These are Go structs that define how records are stored and retrieved using GORM.

## 📁 Contents

- `user.go` — Defines the `User` entity with fields such as email and password hash.

## 🧠 Purpose

The `entities` layer represents the core data structures that map to database tables. These are used throughout the service in:

- 🛠️ Migrations
- 💾 Database operations via repositories
- 🔄 Struct conversions or DTOs

## 🧱 Example

```go
type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
}
```

> This defines a user with a unique email and hashed password, tracked by GORM’s standard model fields (`ID`, `CreatedAt`, etc.).
# 🧾 `repository/` — Data Access Layer

This folder contains the logic for interacting with the database. It abstracts raw database queries behind clean, interface-driven methods.

## 📁 Contents

- `user_repository.go` — Repository for reading user records from the database.

## 🧠 Purpose

Repositories provide a separation of concerns between:
- 🔄 Database operations (SQL/GORM)
- 🧠 Business logic (in controllers)
- 🧪 Testability via interfaces

## 🛠️ Interface

```go
type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}
```

You can inject `UserRepository` into any consumer (e.g., controller) for better testability and flexibility.

## ⚙️ Implementation

```go
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
```
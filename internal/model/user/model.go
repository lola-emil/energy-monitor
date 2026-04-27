package user

import "time"

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	ID int64 `json:"id" db:"id"`

	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password_hash"`

	Name string `json:"name" db:"name"`

	Role UserRole `json:"role" db:"role"`

	IsActive bool `json:"is_active" db:"is_active"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	LastLogin *time.Time `json:"last_login,omitempty" db:"last_login"`
}

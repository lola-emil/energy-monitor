package user

import "time"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`

	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password" json:"-"`
	Role         Role   `db:"user_role" json:"user_role"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// -- DTO
type CreateUserRequest struct {
	Firstname string `json:"firstname" validate:"required,min=2,max=50"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,containsany=0123456789"`
	Role      Role   `json:"user_role" db:"user_role" validate:"required"`
}

type UpdateUserRequest struct {
	Firstname *string `json:"firstname" validate:"omitempty,min=2,max=50"`
	Lastname  *string `json:"lastname"  validate:"omitempty,min=2,max=50"`
	Email     *string `json:"email"     validate:"omitempty,email"`
	Password  *string `json:"password"  validate:"omitempty,min=8"`
	Role      *Role   `json:"user_role" validate:"omitempty,oneof=admin user"`
}

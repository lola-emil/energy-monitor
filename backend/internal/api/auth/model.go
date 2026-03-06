package auth

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Firstname string `json:"firstname" validate:"required,min=2,max=50"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,containsany=0123456789"`
	Role      Role   `json:"user_role" db:"user_role" validate:"required"`
}

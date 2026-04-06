package auth

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,containsany=0123456789"`
}

// internal/auth/errors.go (or internal/domain/errors.go)

package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountDisabled    = errors.New("account disabled")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

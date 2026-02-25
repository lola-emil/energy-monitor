package auth

import (
	"backend/internal/api/user"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) GetUserByEmail(email string) (user.User, error) {
	query := "SELECT * FROM users WHERE email = $1"

	var result user.User
	if err := r.db.Select(&result, query, email); err != nil {
		return user.User{}, err
	}

	return result, nil
}

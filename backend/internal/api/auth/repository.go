package auth

import (
	"backend/internal/api/user"
	"context"

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

func (r *AuthRepo) GetUserByUsername(username string) (user.User, error) {
	query := "SELECT * FROM users WHERE username = $1"

	var result user.User
	if err := r.db.Get(&result, query, username); err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (r *AuthRepo) UserExists(username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)`

	var exists bool
	err := r.db.Get(&exists, query, username)
	return exists, err
}

func (r *AuthRepo) SaveUser(ctx context.Context, record CreateUserRequest) (int64, error) {
	query := `
		INSERT INTO users (
			username,
			password
		) VALUES (
			$1,
			$2
		)

		RETURNING id
	`

	var id int64
	err := r.db.QueryRowContext(ctx, query,
		record.Username,
		record.Password,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

package auth

import (
	"backend/internal/api/user"
	"context"
	"fmt"

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
	if err := r.db.Get(&result, query, email); err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (r *AuthRepo) EmailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`

	var exists bool
	err := r.db.Get(&exists, query, email)
	return exists, err
}

func (r *AuthRepo) SaveUser(ctx context.Context, record CreateUserRequest) (int64, error) {
	query := `
		INSERT INTO users (
			firstname,
			lastname,
			email,
			password,
			user_role
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
	`

	fmt.Printf("%+v\n", record)

	result, err := r.db.ExecContext(ctx, query,
		record.Firstname,
		record.Lastname,
		record.Email,
		record.Password,
		record.Role,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

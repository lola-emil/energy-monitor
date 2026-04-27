package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
	UpdateLastLogin(ctx context.Context, userID int64) error
	Create(ctx context.Context, user *User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*User, error) {
	var user User

	err := r.db.GetContext(ctx, &user, `
		SELECT id, username, password_hash, name, role, is_active, last_login
		FROM users
		WHERE username = $1
	`, username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateLastLogin(ctx context.Context, userID int64) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users
		SET last_login = NOW()
		WHERE id = $1
	`, userID)

	return err
}

func (r *userRepository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password_hash, name, role, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(ctx, query,
		user.Username,
		user.Password,
		user.Name,
		user.Role,
		user.IsActive,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

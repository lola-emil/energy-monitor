package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	var device User
	if err := r.db.Get(&device, query, id); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *UserRepo) GetUsers(ctx context.Context) ([]User, error) {
	query := "SELECT * FROM users WHERE 1"

	var users []User

	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return []User{}, err
	}

	return users, nil
}

func (r *UserRepo) SaveUser(ctx context.Context, record User) (int64, error) {
	query := `
		INSERT INTO users (
			firstname,
			lastname,
			email,
			password,
			user_role
		) VALUES (
			:firstname,
			:lastname,
			:email,
			:password,
			:user_role
		)
	`

	result, err := r.db.ExecContext(ctx, query, map[string]any{
		"firstname": record.Firstname,
		"lastname":  record.Lastname,
		"email":     record.Email,
		"password":  record.PasswordHash,
		"user_role": record.Role,
	})

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *UserRepo) UpdateUserById(ctx context.Context, id int64, data User) (int64, error) {

	query := `
		UPDATE users
		SET
			firstname = :firstname
			lastname = :lastname
			email = :email
			password = :password
			user_role = :user_role
		WHERE
			id = :id
	`

	result, err := r.db.ExecContext(ctx, query, map[string]any{
		"firstname": data.Firstname,
		"lastname":  data.Lastname,
		"email":     data.Email,
		"password":  data.PasswordHash,
		"user_role": data.Role,
	})

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *UserRepo) DeleteUserById(ctx context.Context, id int64) (int64, error) {
	query := "DELETE FROM users WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return 0, nil
	}

	return result.RowsAffected()
}

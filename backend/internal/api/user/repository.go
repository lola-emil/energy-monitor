package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
	query := "SELECT * FROM users"

	var users []User

	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return []User{}, err
	}

	return users, nil
}

func (r *UserRepo) SaveUser(ctx context.Context, record CreateUserRequest) (int64, error) {
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

		RETURNING id
	`

	// fmt.Printf("%+v\n", record)
	var id int64
	err := r.db.QueryRowContext(ctx, query,
		record.Firstname,
		record.Lastname,
		record.Email,
		record.Password,
		record.Role,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) UpdateUserById(ctx context.Context, id int64, data UpdateUserRequest) (int64, error) {
	setClauses := []string{}
	args := []any{}
	argPos := 1

	if data.Firstname != nil {
		setClauses = append(setClauses, fmt.Sprintf("firstname = $%d", argPos))
		args = append(args, data.Firstname)
		argPos++
	}

	if data.Lastname != nil {
		setClauses = append(setClauses, fmt.Sprintf("lastname = $%d", argPos))
		args = append(args, data.Lastname)
		argPos++
	}

	if data.Email != nil {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argPos))
		args = append(args, data.Email)
		argPos++
	}

	if data.Password != nil {
		setClauses = append(setClauses, fmt.Sprintf("password = $%d", argPos))
		args = append(args, data.Password)
		argPos++
	}

	if data.Role != nil {
		setClauses = append(setClauses, fmt.Sprintf("user_role = $%d", argPos))
		args = append(args, data.Role)
		argPos++
	}

	if len(setClauses) == 0 {
		return 0, errors.New("no fields to update")
	}

	query := fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id = $%d
	`, strings.Join(setClauses, ", "), argPos)

	args = append(args, id)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *UserRepo) DeleteUserById(ctx context.Context, id int64) (int64, error) {
	query := "DELETE FROM users WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return 0, nil
	}

	return result.RowsAffected()
}

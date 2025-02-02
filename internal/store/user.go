package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, users *User) error {
	query := `
		INSERT INTO posts (username, email, password)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		users.Username,
		users.Email,
		users.Password,
	).Scan(
		&users.ID,
		&users.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

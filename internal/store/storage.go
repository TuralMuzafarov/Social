package store

import (
	"context"
	"database/sql"
)

type Store struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
	}

	Users interface {
		Create(ctx context.Context, users *User) error
	}
}

func NewStore(db *sql.DB) Store {
	return Store{
		Posts: &PostsStore{db: db},
		Users: &UserStore{db: db},
	}
}

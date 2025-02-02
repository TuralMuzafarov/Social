package store

import "context"

type Store struct {
	Posts interface {
		Create(ctx context.Context, post Post) error
	}

	Users interface {
		Create(ctx context.Context, users User) error
	}
}

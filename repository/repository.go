package repository

import (
	"context"
	"gin/types"

	_ "github.com/lib/pq"
)

type Repository interface {
	SaveUser(ctx context.Context, user *types.User) error
}

type Cache interface {
	CreateUser(ctx context.Context, user *types.User) error
	ReadUser(ctx context.Context) (*types.User, bool)
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateSinger(ctx context.Context, arg CreateSingerParams) (Singer, error)
	GetListSinger(ctx context.Context, arg GetListSingerParams) ([]Singer, error)
	GetMe(ctx context.Context, username string) (GetMeRow, error)
	GetSinger(ctx context.Context, id int64) (Singer, error)
	Login(ctx context.Context, username string) (LoginRow, error)
	Register(ctx context.Context, arg RegisterParams) (User, error)
}

var _ Querier = (*Queries)(nil)

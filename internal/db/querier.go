// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CountSinger(ctx context.Context) (int64, error)
	CreateSinger(ctx context.Context, arg CreateSingerParams) (Singer, error)
	DeleteSinger(ctx context.Context, id int64) error
	GetListSinger(ctx context.Context, arg GetListSingerParams) ([]GetListSingerRow, error)
	GetMe(ctx context.Context, username string) (GetMeRow, error)
	GetSinger(ctx context.Context, id int64) (GetSingerRow, error)
	Login(ctx context.Context, username string) (LoginRow, error)
	Register(ctx context.Context, arg RegisterParams) (User, error)
	UpdateSinger(ctx context.Context, arg UpdateSingerParams) (UpdateSingerRow, error)
}

var _ Querier = (*Queries)(nil)

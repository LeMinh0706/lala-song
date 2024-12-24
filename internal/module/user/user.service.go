package user

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
)

type IUserService interface {
	Register(ctx context.Context, arg db.RegisterParams) (db.User, error)
	Login(ctx context.Context, username, password string) (db.User, error)
}

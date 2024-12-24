//go:build wireinject

package wire

import (
	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/user"
	"github.com/google/wire"
)

func InitUserRouterHandler(q *db.Queries) (*user.UserController, error) {
	wire.Build(
		user.NewUserService,
		user.NewUserController,
	)
	return new(user.UserController), nil
}

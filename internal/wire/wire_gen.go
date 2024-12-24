// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/user"
)

// Injectors from user.wire.go:

func InitUserRouterHandler(q *db.Queries) (*user.UserController, error) {
	iUserService := user.NewUserService(q)
	userController := user.NewUserController(iUserService)
	return userController, nil
}

package initialize

import (
	"database/sql"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/singer"
	"github.com/LeMinh0706/lala-song/internal/module/user"
	"github.com/LeMinh0706/lala-song/util"
)

type ISevice struct {
	UserService   user.IUserService
	SingerService singer.ISingerService
}

func InitService(pg *sql.DB, config util.Config) *ISevice {
	q := db.New(pg)

	userService := user.NewUserService(q)
	singerService := singer.NewSingerService(q)

	return &ISevice{
		UserService:   userService,
		SingerService: singerService,
	}
}

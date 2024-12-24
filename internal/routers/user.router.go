package routers

import (
	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/wire"
	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(f fiber.Router, q *db.Queries) {
	uc, _ := wire.InitUserRouterHandler(q)
	userGroup := f.Group("/users")
	{
		userGroup.Post("/login", uc.Login)
	}
}

package user

import (
	"github.com/LeMinh0706/lala-song/internal/middlewares"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(f fiber.Router, service IUserService, token token.Maker) {
	uc := NewUserController(service, token)
	userGroup := f.Group("/users")
	{
		userGroup.Post("/login", uc.Login)
		userGroup.Post("/register", uc.Register)
	}
	auth := userGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.Get("/me", uc.GetMe)
	}
}

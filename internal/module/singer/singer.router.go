package singer

import (
	"github.com/LeMinh0706/lala-song/internal/middlewares"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

func NewSingerRouter(f fiber.Router, service ISingerService, token token.Maker) {
	sc := NewSingerController(service, token)
	singerGroup := f.Group("/singers")
	{
		singerGroup.Get("/:id", sc.GetSingerById)
		singerGroup.Get("", sc.GetSingers)
	}

	auth := singerGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.Post("", sc.CreateSinger)
		auth.Put("/:id", sc.UpdateSinger)
	}
}

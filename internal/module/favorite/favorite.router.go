package favorite

import (
	"github.com/LeMinh0706/lala-song/internal/middlewares"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

func NewFavoriteRouter(f fiber.Router, service IFavoriteService, token token.Maker) {
	fc := NewFavoriteController(service, token)
	auth := f.Group("/favorite").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.Post("/:id", fc.CreateLikeSong)
		auth.Get("", fc.GetListSongs)
		auth.Get("/:id", fc.GetLikeSong)
		auth.Delete("/:id", fc.UnlikeSong)
	}
}

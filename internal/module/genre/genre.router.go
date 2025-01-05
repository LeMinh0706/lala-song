package genre

import (
	"github.com/LeMinh0706/lala-song/internal/middlewares"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

func NewGenreRouter(f fiber.Router, service IGenreService, token token.Maker) {
	gc := NewGenreController(service)
	genreGroup := f.Group("/genres")
	{

	}
	auth := genreGroup.Group("").Use(middlewares.AuthorizeAdminMiddleware(token))
	{
		auth.Post("", gc.CreateGenre)
	}
}

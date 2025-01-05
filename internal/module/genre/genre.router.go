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
		genreGroup.Get("/:id", gc.GetGenreById)
		genreGroup.Get("", gc.GetGenres)
	}
	auth := genreGroup.Group("").Use(middlewares.AuthorizeAdminMiddleware(token))
	{
		auth.Post("", gc.CreateGenre)
		auth.Put("/:id", gc.UpdateGenre)
		auth.Delete("/:id", gc.DeleteGenre)
	}
}

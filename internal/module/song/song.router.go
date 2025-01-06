package song

import (
	"github.com/LeMinh0706/lala-song/internal/middlewares"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

func NewSongRouter(f fiber.Router, service ISongService, token token.Maker) {
	sc := NewSongController(service)
	songGroup := f.Group("/songs")
	{
		songGroup.Get("/:id", sc.GetSongById)
	}
	auth := songGroup.Group("").Use(middlewares.AuthorizeAdminMiddleware(token))
	{
		auth.Post("", sc.CreateSong)
		auth.Post("/feature", sc.AddFeatureSong)
		auth.Post("/genre", sc.AddGenre)
	}
}

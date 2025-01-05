package album

import (
	"github.com/LeMinh0706/lala-song/internal/middlewares"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

func NewAlbumRouter(f fiber.Router, service IAlbumService, token token.Maker) {
	ac := NewAlBumController(service)
	albumGroup := f.Group("/albums")
	{
		albumGroup.Get("", ac.GetAlbums)
	}
	auth := albumGroup.Group("").Use(middlewares.AuthorizeAdminMiddleware(token))
	{
		auth.Post("", ac.CreateAlbum)
		auth.Put("/:id", ac.UpdateAlbum)
		auth.Post("/soft/:id", ac.DeleteAlbum)
	}
}

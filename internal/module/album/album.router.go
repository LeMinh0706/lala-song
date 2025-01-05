package album

import (
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

func NewAlbumRouter(f fiber.Router, service IAlbumService, token token.Maker) {
	ac := NewAlBumController(service)
	albumGroup := f.Group("/albums")
	{

	}
	{
		albumGroup.Post("", ac.CreateAlbum)
	}
}

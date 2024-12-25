package server

import (
	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/routers"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewRouter() {
	q := db.New(s.DBConn)
	a := s.Router.Group("/api")
	{
		a.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})
		routers.NewUserRouter(a, q, s.TokenMaker)
	}

}

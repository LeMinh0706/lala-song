package server

import "github.com/gofiber/fiber/v2"

func (s *Server) NewRouter() {
	a := s.Router.Group("/api")
	{
		a.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})
	}
}

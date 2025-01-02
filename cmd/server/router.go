package server

import (
	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/routers"
	_ "github.com/LeMinh0706/lala-song/swag/docs"
	"github.com/gofiber/swagger"
)

func (s *Server) NewRouter() {
	q := db.New(s.DBConn)

	s.Router.Get("/swagger/*", swagger.HandlerDefault)
	a := s.Router.Group("/api")
	{
		routers.NewUserRouter(a, q, s.TokenMaker)
	}
}

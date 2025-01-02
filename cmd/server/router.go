package server

import (
	"github.com/LeMinh0706/lala-song/internal/initialize"
	"github.com/LeMinh0706/lala-song/internal/module/user"
	_ "github.com/LeMinh0706/lala-song/swag/docs"
	"github.com/gofiber/swagger"
)

func (s *Server) NewRouter() {

	initService := initialize.InitService(s.DBConn, s.Config)

	s.Router.Get("/swagger/*", swagger.HandlerDefault)
	a := s.Router.Group("/api")
	{
		user.NewUserRouter(a, initService.UserService, s.TokenMaker)
	}
}

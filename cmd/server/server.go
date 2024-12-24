package server

import (
	"database/sql"

	"github.com/LeMinh0706/lala-song/util"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Config util.Config
	Router *fiber.App
	DBConn *sql.DB
}

func NewServer(pg *sql.DB, config util.Config) (*Server, error) {
	app := fiber.New(fiber.Config{
		Prefork:   true,
		BodyLimit: 6 * 1024 * 1024,
	})
	server := &Server{
		Config: config,
		Router: app,
		DBConn: pg,
	}
	server.NewRouter()
	server.Static()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Listen(server.Config.ServerAddress)
}

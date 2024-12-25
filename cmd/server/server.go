package server

import (
	"database/sql"
	"log"

	"github.com/LeMinh0706/lala-song/token"
	"github.com/LeMinh0706/lala-song/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Config     util.Config
	Router     *fiber.App
	DBConn     *sql.DB
	TokenMaker token.Maker
}

func NewServer(pg *sql.DB, config util.Config) (*Server, error) {
	token, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		log.Fatal("can't create token")
	}
	app := fiber.New(fiber.Config{
		Prefork:   true,
		BodyLimit: 6 * 1024 * 1024,
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	server := &Server{
		Config:     config,
		Router:     app,
		DBConn:     pg,
		TokenMaker: token,
	}
	server.NewRouter()
	server.Static()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Listen(server.Config.ServerAddress)
}

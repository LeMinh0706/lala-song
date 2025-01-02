package main

import (
	"log"

	"github.com/LeMinh0706/lala-song/cmd/server"
	"github.com/LeMinh0706/lala-song/internal/initialize"
	"github.com/LeMinh0706/lala-song/util"
	_ "github.com/lib/pq"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8070
// @BasePath /api
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config:", err)
	}
	pg, err := initialize.Postgres(config)
	if err != nil {
		log.Fatal("Cannot postgres: ", err)
	}
	server, err := server.NewServer(pg, config)
	if err != nil {
		log.Fatal("Cannot load server: ", err)
	}

	server.Start(config.ServerAddress)
}

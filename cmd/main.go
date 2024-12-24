package main

import (
	"log"

	"github.com/LeMinh0706/lala-song/cmd/server"
	"github.com/LeMinh0706/lala-song/internal/initialize"
	"github.com/LeMinh0706/lala-song/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config:", err)
	}
	pg, err := initialize.Postgres(config)
	defer pg.Close()
	server, err := server.NewServer(pg, config)
	if err != nil {
		log.Fatal("Cannot load server: ", err)
	}
	server.Start(config.ServerAddress)
}

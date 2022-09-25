package main

import (
	"log"

	"github.com/sirijagadeesh/playrestapi1/config"
	"github.com/sirijagadeesh/playrestapi1/server"
)

func main() {
	log.Println("Hello Ramya")

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Println(err)

		return
	}

	server.Start(cfg)
}

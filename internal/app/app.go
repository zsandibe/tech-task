package app

import (
	"doodocsTask/config"
	"doodocsTask/internal/delivery"
	"doodocsTask/internal/server"
	"doodocsTask/internal/service"
	"log"
)

func Start() {
	config, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	service := service.NewService()

	delivery := delivery.NewHandler(service)

	server := new(server.Server)

	if err := server.Run(config, delivery.Routes()); err != nil {
		log.Fatalf("Error running server: %s\n", err)
	}
}

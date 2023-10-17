package main

import (
	"doodocsTask/config"
	"doodocsTask/internal/delivery"
	"doodocsTask/internal/server"
	"doodocsTask/internal/service"
	"log"
)

func main() {
	config, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	service := service.NewService(config)

	delivery := delivery.NewHandler(service)

	server := new(server.Server)

	if err := server.Run(config, delivery.Routes()); err != nil {
		log.Fatalf("Error running server: %s\n", err)
	}
}

package main

import (
	"Internship_backend_avito/configs"
	v1 "Internship_backend_avito/internal/controller/http/v1"
	"Internship_backend_avito/internal/repository/postgresdb"
	"Internship_backend_avito/internal/service"
	"Internship_backend_avito/pkg/httpserver"
	"log"
)

func main() {
	cfg := configs.MustLoad()

	repos := postgresdb.NewRepository()
	serv := service.NewService(repos)
	handlers := v1.NewHandler(serv)

	server := new(httpserver.Server)
	if err := server.Run(cfg.App.Port, handlers.InitHandler()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

package main

import (
	v1 "Internship_backend_avito/internal/controller/http/v1"
	"Internship_backend_avito/internal/repository/postgresdb"
	"Internship_backend_avito/internal/service"
	"Internship_backend_avito/pkg/httpserver"
	"log"
)

func main() {
	repos := postgresdb.NewRepository()
	serv := service.NewService(repos)
	handlers := v1.NewHandler(serv)

	server := new(httpserver.Server)
	if err := server.Run("8000", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}

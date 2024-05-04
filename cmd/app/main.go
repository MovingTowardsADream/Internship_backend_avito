package main

import (
	v1 "Internship_backend_avito/internal/controller/http/v1"
	"Internship_backend_avito/pkg/httpserver"
	"log"
)

func main() {
	handlers := new(v1.Handler)
	server := new(httpserver.Server)
	if err := server.Run("8000", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}

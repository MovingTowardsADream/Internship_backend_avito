package main

import (
	"Internship_backend_avito/configs"
	v1 "Internship_backend_avito/internal/controller/http/v1"
	"Internship_backend_avito/internal/repository/postgresdb"
	"Internship_backend_avito/internal/service"
	"Internship_backend_avito/pkg/httpserver"
	"Internship_backend_avito/pkg/postgres"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	cfg := configs.MustLoad()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	logrus.Info("Initializing postgres...")

	// TODO Env + Config
	db, err := postgres.NewPostgresDB(
		"postgres://postgres:admin@localhost:5432/postgres",
		postgres.MaxPoolSize(20))

	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %w", err))
	}
	defer db.Close()

	// TODO ping db

	// Repositories
	logrus.Info("Initializing repositories...")

	repos := postgresdb.NewRepository(db)
	serv := service.NewService(repos)
	handlers := v1.NewHandler(serv)

	server := new(httpserver.Server)
	if err = server.Run(cfg.App.Port, handlers.InitHandler()); err != nil {
		logrus.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

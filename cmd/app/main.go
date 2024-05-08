package main

import (
	"Internship_backend_avito/configs"
	v1 "Internship_backend_avito/internal/controller/http/v1"
	"Internship_backend_avito/internal/repository/postgresdb"
	"Internship_backend_avito/internal/service"
	"Internship_backend_avito/pkg/httpserver"
	"Internship_backend_avito/pkg/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	cfg := configs.MustLoad()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.ConfigDB{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Username: cfg.Database.Username,
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	})

	if err != nil {
		logrus.Fatalf("Failed initialization database: %s", err.Error())
	}

	repos := postgresdb.NewRepository(db)
	serv := service.NewService(repos)
	handlers := v1.NewHandler(serv)

	server := new(httpserver.Server)
	if err := server.Run(cfg.App.Port, handlers.InitHandler()); err != nil {
		logrus.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

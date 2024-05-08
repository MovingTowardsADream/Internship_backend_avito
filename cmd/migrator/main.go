package main

import (
	"Internship_backend_avito/configs"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {

	var migrationsPath string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.Parse()

	if migrationsPath == "" {
		panic("Migrations-path is required")
	}

	cfg := configs.MustLoad()

	if err := godotenv.Load(); err != nil {
		panic("Failed reading db password")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("postgres://%s:admin@%s:%s/%s?sslmode=%s", cfg.Database.Username, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("Migrations applied")
}

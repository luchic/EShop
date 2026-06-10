package main

import (
	"log"
	"shop/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.NewConfig("./config.json")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	m, err := migrate.New(
		"file://./migrations",
		cfg.ConnectionString)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("Migrations applied successfully!")
}

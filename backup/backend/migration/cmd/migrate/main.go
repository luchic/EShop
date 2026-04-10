package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := os.Getenv("AUTH_DATABASE_URL")
	if connectionString == "" {
		log.Fatal("AUTH_DATABASE_URL is required")
	}

	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	if command == "" {
		log.Fatal("usage: go run ./cmd/migrate/main.go [up|down]")
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("unable to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("unable to create migrate instance: %v", err)
	}

	switch command {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migration up failed: %v", err)
		}
		fmt.Println("migration up completed")
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migration down failed: %v", err)
		}
		fmt.Println("migration down completed")
	default:
		log.Fatalf("unknown command %q", command)
	}
}

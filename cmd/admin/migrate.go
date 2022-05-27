package main

import (
	"log"

	"github.com/ivan-sabo/tic-tac-toe/internal/infrastructure/postgres"
	_ "github.com/lib/pq" // Register the postgres database/sql driver
)

func main() {
	// Setup dependencies
	db, err := postgres.Open(postgres.Config{
		Host:       "localhost",
		User:       "postgres",
		Password:   "postgres",
		DisableTLS: true,
		Name:       "postgres",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := postgres.Migrate(db); err != nil {
		log.Fatal("applying migrations: ", err)
	}
	log.Println("Migrations complete")
}

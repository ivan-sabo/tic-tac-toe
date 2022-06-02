package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/ivan-sabo/tic-tac-toe/cmd/api/internal/handlers"
	"github.com/ivan-sabo/tic-tac-toe/internal/infrastructure/postgres"
	"github.com/ivan-sabo/tic-tac-toe/internal/infrastructure/postgres/repository"
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

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gameRepo := repository.GamePostgre{DB: db}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers.GetRouter(&gameRepo),
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

package main

import (
	"context"
	"cryptcloud-core/internal/config"
	"cryptcloud-core/internal/db"
	"cryptcloud/internal/otellogging"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// main is the entrypoint for the application.
//
// It loads the configuration file, sets up the database connection, and starts
// the server.
//
// Parameters:
//
//	err error - error from loading configuration file.
//
// Returns:
//
//	error - if there was an error running the application.
func main() {
	if err := run(nil); err != nil {
		log.Fatalf("error running application: %v", err)
	}
}

// run sets up the application and starts the server.
//
// Parameters:
//
//	err *error - error from loading configuration file.
//
// Returns:
//
//	error - if there was an error running the application.
func run(err *error) (retErr error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// load config
	config, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading configuration file: %w", err)
	}

	// Set up database.
	if err = db.GetDatabase(ctx, config); err != nil {
		return fmt.Errorf("error setting up database: %w", err)
	}

	// Set up OpenTelemetry.

	// Set up OpenTelemetry.

	otelShutdown := otellogging.SetupOTelSDK(ctx) // Handle shutdown properly so nothing leaks.
	if err != nil {
		return fmt.Errorf("error setting up OpenTelemetry: %w", err)
	} // Handle shutdown properly so nothing leaks.
	defer func() {
		retErr = errors.Join(retErr, otelShutdown(context.Background()))
	}()

	// Set up router.
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	return nil
}

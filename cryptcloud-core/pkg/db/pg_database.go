package db

import (
	"context"
	"cryptcloud/pkg/config"
	"fmt"
	"log"

	"entgo.io/ent"

	_ "github.com/lib/pq"
)

func GetDatabase(ctx context.Context) *ent.Client {
	ctx = context.Background()
	// Load config
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Set up database
	databasClient, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabaseName, config.DatabasePassword))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer databasClient.Close()

	// Run the auto migration tool.
	if err := databasClient.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return databasClient

}

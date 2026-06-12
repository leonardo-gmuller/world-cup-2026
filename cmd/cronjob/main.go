package main

import (
	"context"
	"log"
	"os"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
	api_football "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api_football"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/cronjob"
	football_data "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/football_data"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/logger"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/service/hash"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/service/jwt"
)

var Version = "dev"

func main() {
	ctx := context.Background()
	logger.Init()

	cfg, err := config.New(Version)
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}

	postgresClient, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to start postgres: %v", err)
	}
	defer postgresClient.Close()

	hashService := hash.New()
	jwtService := jwt.New(cfg.JWT.Secret)
	footballClient := football_data.New(cfg)
	liveScoreClient := api_football.New(cfg)

	appl := app.New(
		ctx,
		cfg,
		postgresClient,
		hashService,
		jwtService,
		footballClient,
		liveScoreClient,
	)

	usecase := appl.NewJobControlUseCase(postgresClient.Pool)

	cronApp := cronjob.New(cfg, usecase)

	if err := cronApp.Run(os.Args); err != nil {
		log.Fatalf("failed to run cronjob: %v", err)
	}
}

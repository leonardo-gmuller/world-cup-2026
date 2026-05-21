package main

import (
	"context"
	"log"
	"os"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/cronjob"
	football_api "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/football_api"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/redis"
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

	redisClient, err := redis.NewClient(cfg.Redis)
	if err != nil {
		log.Fatalf("failed to start redis: %v", err)
	}

	hashService := hash.New()
	jwtService := jwt.New(cfg.JWT.Secret)
	footballClient := football_api.New(cfg)

	appl := app.New(
		ctx,
		cfg,
		postgresClient,
		redisClient,
		hashService,
		jwtService,
		footballClient,
	)

	usecase := appl.NewJobControlUseCase(postgresClient.Pool)

	cronApp := cronjob.New(cfg, usecase)

	if err := cronApp.Run(os.Args); err != nil {
		log.Fatalf("failed to run cronjob: %v", err)
	}
}

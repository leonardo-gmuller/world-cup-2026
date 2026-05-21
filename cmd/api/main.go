package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api"
	football_api "github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/football_api"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/redis"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/logger"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/service/hash"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/service/jwt"
	"golang.org/x/sync/errgroup"
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

	stopCtx, stop := signal.NotifyContext(
		ctx,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	defer stop()

	group, groupCtx := errgroup.WithContext(stopCtx)

	server := &http.Server{
		Addr: cfg.Server.Address,
		BaseContext: func(_ net.Listener) context.Context {
			return stopCtx
		},
		Handler:      api.New(cfg, appl).Handler,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	group.Go(func() error {
		log.Printf("starting api server on %s", cfg.Server.Address)

		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	group.Go(func() error {
		<-groupCtx.Done()

		log.Printf("stopping api; interrupt signal received")

		timeoutCtx, cancel := context.WithTimeout(
			context.Background(),
			cfg.App.GracefulShutdownTimeout,
		)
		defer cancel()

		var errs error

		if err := server.Shutdown(timeoutCtx); err != nil {
			errs = errors.Join(
				errs,
				fmt.Errorf("failed to stop server: %w", err),
			)
		}

		return errs
	})

	if err := group.Wait(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("api exit reason: %v", err)
	}
}

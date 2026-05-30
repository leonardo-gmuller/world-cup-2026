package handler

import (
	"github.com/cep21/circuit/v3"
	"github.com/cep21/circuit/v3/closers/hystrix"
	"github.com/go-chi/chi/v5"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
)

type Handler struct {
	circuitManager *circuit.Manager
	cfg            config.Config
	app            *app.App
}

func New(cfg config.Config, a *app.App) Handler {
	hystrixFactory := hystrix.Factory{
		ConfigureOpener: hystrix.ConfigureOpener{
			ErrorThresholdPercentage: int64(cfg.CircuitBreaker.ErrorPercentThreshold),
			RequestVolumeThreshold:   int64(cfg.CircuitBreaker.RequestVolumeThreshold),
		},
		ConfigureCloser: hystrix.ConfigureCloser{
			SleepWindow: cfg.CircuitBreaker.SleepWindow,
		},
	}

	defaultFactory := func(_ string) circuit.Config {
		return circuit.Config{
			Execution: circuit.ExecutionConfig{
				MaxConcurrentRequests: int64(cfg.CircuitBreaker.MaxConcurrentRequests),
				Timeout:               cfg.CircuitBreaker.Timeout,
			},
		}
	}

	circuitManager := &circuit.Manager{
		DefaultCircuitProperties: []circuit.CommandPropertiesConstructor{
			defaultFactory,
			hystrixFactory.Configure,
		},
	}

	return Handler{
		circuitManager: circuitManager,
		cfg:            cfg,
		app:            a,
	}
}

func RegisterBasicRoutes(router chi.Router, cfg config.Config, a *app.App) {
	h := New(cfg, a)

	h.healthcheckSetup(router)
}

func RegisterAuthRoutes(
	router chi.Router,
	cfg config.Config,
	a *app.App,
) {
	h := New(cfg, a)

	h.authSetupRoutes(router)
}

func RegisterPrivateRoutes(
	router chi.Router,
	cfg config.Config,
	a *app.App,
) {
	h := New(cfg, a)

	h.userSetupRoutes(router)
	h.groupSetupRoutes(router)
	h.matchSetupRoutes(router)
	h.predictionSetupRoutes(router)
	h.rankingSetupRoutes(router)
	h.dashboardSetupRoutes(router)
}

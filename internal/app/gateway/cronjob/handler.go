package cronjob

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

type useCase interface {
	ImportMatches(ctx context.Context) error
	CreateJobsControl(ctx context.Context, jobID types.Job) error
	UpdateJobsControl(ctx context.Context, jobID types.Job) error
	CalculateMatchPredictions(ctx context.Context) error
}

type Handler struct {
	cfg     config.Config
	useCase useCase
}

func NewHandler(cfg config.Config, uc useCase) *Handler {
	handler := Handler{
		cfg:     cfg,
		useCase: uc,
	}

	return &handler
}

package cronjob

import (
	"context"
	"log/slog"
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/scheduler"
)

func (h *Handler) StartScheduler(ctx context.Context) error {
	scheduler, err := scheduler.New(ctx)
	if err != nil {
		return err
	}

	// Import/update matches every 30min
	err = scheduler.Every(30*time.Minute, func(ctx context.Context) error {
		return h.RunJob(
			ctx,
			types.ImportMatches,
			h.ImportMatches,
		)
	})
	if err != nil {
		return err
	}

	// Update live scores/status every 5min
	err = scheduler.Every(5*time.Minute, func(ctx context.Context) error {
		return h.RunJob(
			ctx,
			types.SyncLiveResults,
			h.SyncLiveResults,
		)
	})
	if err != nil {
		return err
	}

	// Calculate prediction points every 2min
	err = scheduler.Every(2*time.Minute, func(ctx context.Context) error {
		return h.RunJob(
			ctx,
			types.CalculatePredictionPoints,
			h.CalculateMatchPredictions,
		)
	})
	if err != nil {
		return err
	}

	scheduler.Start()

	slog.InfoContext(ctx, "[scheduler] started")

	<-ctx.Done()

	return scheduler.Stop()
}

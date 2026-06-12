package cronjob

import (
	"context"
	"log/slog"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
	"github.com/urfave/cli/v2"
)

func New(cfg config.Config, useCase useCase) *cli.App {
	handler := NewHandler(cfg, useCase)

	return &cli.App{
		Commands: []*cli.Command{
			{
				Name:  string(types.ImportMatches),
				Usage: "Import matches from external source",
				Action: func(ctx *cli.Context) error {
					return handler.RunJob(ctx.Context, types.ImportMatches, handler.ImportMatches)
				},
			},
			{
				Name:  string(types.CalculatePredictionPoints),
				Usage: "Calculate match predictions points",
				Action: func(ctx *cli.Context) error {
					return handler.RunJob(ctx.Context, types.CalculatePredictionPoints, handler.CalculateMatchPredictions)
				},
			},
			{
				Name:  "start-scheduler",
				Usage: "Start the cronjob scheduler to run periodic tasks",
				Action: func(ctx *cli.Context) error {
					return handler.StartScheduler(ctx.Context)
				},
			},
			{
				Name:  "sync-live-results",
				Usage: "Sync live match results from external source",
				Action: func(ctx *cli.Context) error {
					return handler.RunJob(ctx.Context, types.SyncLiveResults, handler.SyncLiveResults)
				},
			},
		},
	}
}

func (h *Handler) RunJob(ctx context.Context, jobID types.Job, action func(context.Context) error) error {
	if err := h.useCase.CreateJobsControl(ctx, jobID); err != nil {
		return err
	}

	if err := action(ctx); err != nil {
		slog.ErrorContext(ctx, err.Error())
		return err
	}

	return h.useCase.UpdateJobsControl(ctx, jobID)
}

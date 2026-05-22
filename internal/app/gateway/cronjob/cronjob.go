package cronjob

import (
	"fmt"
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
				Action: runJobAction(func(ctx *cli.Context) error {
					return handler.ImportMatches(ctx.Context)
				}, handler, types.ImportMatches),
			},
			{
				Name:  string(types.CalculatePredictionPoints),
				Usage: "Calculate match predictions points",
				Action: runJobAction(func(ctx *cli.Context) error {
					return handler.CalculateMatchPredictions(ctx.Context)
				}, handler, types.CalculatePredictionPoints),
			},
		},
	}
}

func runJobAction(action cli.ActionFunc, handler *Handler, jobID types.Job) cli.ActionFunc {
	return func(cliCtx *cli.Context) error {
		err := handler.useCase.CreateJobsControl(cliCtx.Context, jobID)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		err = action(cliCtx)
		if err != nil {
			slog.ErrorContext(cliCtx.Context, err.Error())

			return fmt.Errorf("%w", err)
		}

		err = handler.useCase.UpdateJobsControl(cliCtx.Context, jobID)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		return nil
	}
}

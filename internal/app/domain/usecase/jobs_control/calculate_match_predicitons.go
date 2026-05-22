package jobcontrol_usecase

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

func (u *JobControlUsecase) CalculateMatchPredictions(ctx context.Context) error {
	slog.InfoContext(ctx, "[predictions] starting prediction calculation")

	matches, err := u.matchUsecase.CalculateFinishedMatches(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "[predictions] failed fetching finished matches", "error", err)
		return err
	}

	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(5)

	for _, match := range matches {
		match := match
		group.Go(func() error {
			slog.InfoContext(ctx, "[predictions] calculating match", "match_id", match.ID)

			return u.predictionUsecase.CalculateMatchPredictions(ctx, match.ID)
		})
	}

	if err := group.Wait(); err != nil {
		slog.ErrorContext(ctx, "[predictions] calculation failed", "error", err)
		return err
	}

	slog.InfoContext(ctx, "[predictions] prediction calculation finished successfully")
	return nil
}

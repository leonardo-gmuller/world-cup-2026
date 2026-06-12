package jobcontrol_usecase

import (
	"context"
	"log/slog"
	"time"
)

func (u *JobControlUsecase) ImportMatches(ctx context.Context) error {
	slog.InfoContext(ctx, "[cronjob] starting match import")

	shouldImport, err := u.ShouldImportMatches(
		ctx,
		30*time.Minute,
	)
	if err != nil {
		slog.ErrorContext(ctx, "[cronjob] failed checking import cooldown", "error", err)
		return err
	}

	if !shouldImport {
		slog.InfoContext(
			ctx,
			"[cronjob] skipping import",
			"reason", "cooldown_active",
		)
		return nil
	}

	slog.InfoContext(ctx, "[cronjob] importing matches from football-data")

	if err := u.matchUsecase.ImportMatches(ctx); err != nil {
		slog.ErrorContext(ctx, "[cronjob] failed importing matches", "error", err)
		return err
	}

	slog.InfoContext(ctx, "[cronjob] match import finished successfully")

	return nil
}

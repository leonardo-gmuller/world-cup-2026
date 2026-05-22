package jobcontrol_usecase

import (
	"context"
	"log/slog"
	"time"
)

func (u *JobControlUsecase) ImportMatches(ctx context.Context) error {
	slog.InfoContext(ctx, "[cronjob] starting match import")

	hasLive, err := u.matchUsecase.HasLiveMatches(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "[cronjob] failed checking live matches", "error", err)
		return err
	}

	shouldImport := hasLive

	slog.InfoContext(
		ctx,
		"[cronjob] live match status checked",
		"has_live_match", hasLive,
	)

	if !hasLive {
		shouldImport, err = u.ShouldImportMatches(
			ctx,
			15*time.Minute,
		)
		if err != nil {
			slog.ErrorContext(ctx, "[cronjob] failed checking import cooldown", "error", err)
			return err
		}

		slog.InfoContext(
			ctx,
			"[cronjob] cooldown validation finished",
			"should_import", shouldImport,
		)
	}

	if !shouldImport {
		slog.InfoContext(
			ctx,
			"[cronjob] skipping import",
			"reason", "cooldown_active_and_no_live_match",
		)
		return nil
	}

	slog.InfoContext(ctx, "[cronjob] importing matches from external api")

	if err := u.matchUsecase.ImportMatches(ctx); err != nil {
		slog.ErrorContext(ctx, "[cronjob] failed importing matches", "error", err)

		return err
	}

	slog.InfoContext(ctx, "[cronjob] match import finished successfully")

	return nil
}

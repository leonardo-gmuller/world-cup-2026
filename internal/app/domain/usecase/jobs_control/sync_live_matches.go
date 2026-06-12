package jobcontrol_usecase

import (
	"context"
	"log/slog"
)

func (u *JobControlUsecase) SyncLiveResults(ctx context.Context) error {
	slog.InfoContext(ctx, "[cronjob] starting live results sync")

	if err := u.matchUsecase.SyncLiveResults(ctx); err != nil {
		slog.ErrorContext(ctx, "[cronjob] failed syncing live results", "error", err)
		return err
	}

	slog.InfoContext(ctx, "[cronjob] live results sync finished successfully")

	return nil
}

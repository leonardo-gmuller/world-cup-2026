package jobcontrol_usecase

import (
	"context"
	"time"
)

func (u *JobControlUsecase) ImportMatches(ctx context.Context) error {
	hasLive, err := u.matchUsecase.HasLiveMatches(ctx)
	if err != nil {
		return err
	}

	shouldImport := hasLive

	if !hasLive {
		shouldImport, err = u.ShouldImportMatches(
			ctx,
			15*time.Minute,
		)
		if err != nil {
			return err
		}
	}

	if !shouldImport {
		return nil
	}

	if err := u.matchUsecase.ImportMatches(ctx); err != nil {
		return err
	}

	return nil
}

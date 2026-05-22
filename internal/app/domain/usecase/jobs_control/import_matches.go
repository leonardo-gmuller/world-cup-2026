package jobcontrol_usecase

import "context"

func (u *JobControlUsecase) ImportMatches(ctx context.Context) error {
	if err := u.matchUsecase.ImportMatches(ctx); err != nil {
		return err
	}

	return nil
}

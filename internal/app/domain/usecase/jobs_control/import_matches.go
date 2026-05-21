package jobcontrol_usecase

import "context"

func (u *JobControlUsecase) ImportMatches(ctx context.Context) error {
	if err := u.matchUsecase.ImportMatches(ctx); err != nil {
		return err
	}

	matches, err := u.matchUsecase.CalculateFinishedMatches(ctx)
	if err != nil {
		return err
	}

	for _, match := range matches {
		if err := u.predictionUsecase.CalculateMatchPredictions(ctx, match.ID); err != nil {
			return err
		}
	}

	return nil
}

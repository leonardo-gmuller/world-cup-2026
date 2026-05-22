package jobcontrol_usecase

import "context"

func (u *JobControlUsecase) CalculateMatchPredictions(ctx context.Context) error {
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

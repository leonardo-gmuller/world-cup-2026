package prediction_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/constants"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/erring"
)

func (u *PredictionUseCase) CalculateMatchPredictions(
	ctx context.Context,
	matchID int64,
) error {
	match, err := u.matchRepository.GetMatchByID(ctx, matchID)
	if err != nil {
		return err
	}

	if match.HomeScore == nil || match.AwayScore == nil {
		return erring.ErrMatchWithoutScore
	}

	if match.Status != constants.MatchStatusFinished {
		return erring.ErrMatchNotFinished
	}

	stageWeight, err := u.predictionRepo.GetStageWeight(ctx, match.Stage)
	if err != nil {
		return err
	}

	predictions, err := u.predictionRepo.ListPredictionsByMatch(ctx, matchID)
	if err != nil {
		return err
	}

	for _, prediction := range predictions {
		points := calculatePredictionPoints(
			prediction.HomeScore,
			prediction.AwayScore,
			*match.HomeScore,
			*match.AwayScore,
			stageWeight.Weight,
		)

		err = u.predictionRepo.UpdatePredictionPoints(ctx, prediction.ID, points)
		if err != nil {
			return err
		}
	}

	return nil
}

func calculatePredictionPoints(
	predictedHome int,
	predictedAway int,
	realHome int,
	realAway int,
	weight float64,
) float64 {
	basePoints := 0.0

	if predictedHome == realHome && predictedAway == realAway {
		basePoints = 5
	} else if getResult(predictedHome, predictedAway) == getResult(realHome, realAway) {
		basePoints = 3
	}

	return basePoints * weight
}

func getResult(homeScore int, awayScore int) string {
	if homeScore > awayScore {
		return "home"
	}

	if awayScore > homeScore {
		return "away"
	}

	return "draw"
}

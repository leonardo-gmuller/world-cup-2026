package prediction_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *PredictionUseCase) GetPredictionByMatchAndUser(
	ctx context.Context,
	groupID int64,
	userID int64,
	matchID int64,
) (*entity.Prediction, error) {
	return u.predictionRepo.GetPrediction(ctx, groupID, userID, matchID)
}

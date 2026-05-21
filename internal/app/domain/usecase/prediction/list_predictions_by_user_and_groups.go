package prediction_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *PredictionUseCase) ListPredictionsByUserAndGroup(
	ctx context.Context,
	groupID int64,
	userID int64,
) ([]entity.Prediction, error) {
	return u.predictionRepo.ListPredictionsByUserAndGroup(ctx, groupID, userID)
}

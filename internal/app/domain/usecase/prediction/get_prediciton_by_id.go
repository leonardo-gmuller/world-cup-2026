package prediction_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *PredictionUseCase) GetPredictionByID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Prediction, error) {
	return u.predictionRepo.GetPredictionByID(ctx, id)
}

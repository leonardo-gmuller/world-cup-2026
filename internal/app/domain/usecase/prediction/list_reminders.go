package prediction_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/dto"
)

func (u *PredictionUseCase) ListPredictionRemindersByUserID(
	ctx context.Context,
	userID int64,
) ([]dto.PredictionReminderOutput, error) {

	return u.predictionRepo.ListPredictionRemindersByUserID(ctx, userID)
}

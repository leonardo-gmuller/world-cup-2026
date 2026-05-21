package prediction_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *PredictionUseCase) ListPredictionsByUserAndGroup(
	ctx context.Context,
	groupID string,
	userID int64,
) ([]entity.Prediction, error) {
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		return nil, err
	}

	group, err := u.groupRepository.GetGroupByID(ctx, groupUUID)
	if err != nil {
		return nil, err
	}
	return u.predictionRepo.ListPredictionsByUserAndGroup(ctx, group.ID, userID)
}

package prediction_usecase

import (
	"context"

	"github.com/google/uuid"
)

type RankingItemOutput struct {
	UserID           int64
	UserUUID         uuid.UUID
	Name             string
	TotalPoints      float64
	PredictionsCount int64
}

func (u *PredictionUseCase) GetGroupRanking(
	ctx context.Context,
	groupID string,
) ([]RankingItemOutput, error) {
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		return nil, err
	}

	group, err := u.groupRepository.GetGroupByID(ctx, groupUUID)
	if err != nil {
		return nil, err
	}

	return u.predictionRepo.GetGroupRanking(ctx, group.ID)
}

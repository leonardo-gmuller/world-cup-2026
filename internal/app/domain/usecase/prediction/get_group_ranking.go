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
	groupID int64,
) ([]RankingItemOutput, error) {
	return u.predictionRepo.GetGroupRanking(ctx, groupID)
}

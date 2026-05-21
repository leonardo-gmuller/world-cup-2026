package prediction_repository

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func (r *PredictionRepository) UpsertPrediction(
	ctx context.Context,
	prediction entity.Prediction,
) (*entity.Prediction, error) {
	row, err := r.Queries.UpsertPrediction(ctx, sqlc.UpsertPredictionParams{
		Uuid:      prediction.UUID,
		GroupID:   prediction.GroupID,
		UserID:    prediction.UserID,
		MatchID:   prediction.MatchID,
		HomeScore: int32(prediction.HomeScore),
		AwayScore: int32(prediction.AwayScore),
	})
	if err != nil {
		return nil, err
	}

	return mapPrediction(row), nil
}

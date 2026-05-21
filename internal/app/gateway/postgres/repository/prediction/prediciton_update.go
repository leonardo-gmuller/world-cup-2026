package prediction_repository

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func (r *PredictionRepository) UpdatePredictionPoints(
	ctx context.Context,
	predictionID int64,
	points float64,
) error {
	return r.Queries.UpdatePredictionPoints(ctx, sqlc.UpdatePredictionPointsParams{
		ID:     predictionID,
		Points: points,
	})
}

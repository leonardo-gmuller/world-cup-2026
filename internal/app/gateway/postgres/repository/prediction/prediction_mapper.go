package prediction_repository

import (
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func mapPrediction(row sqlc.Prediction) *entity.Prediction {
	return &entity.Prediction{
		ID:         row.ID,
		UUID:       row.Uuid,
		GroupID:    row.GroupID,
		UserID:     row.UserID,
		MatchID:    row.MatchID,
		HomeScore:  int(row.HomeScore),
		AwayScore:  int(row.AwayScore),
		Points:     row.Points,
		Calculated: row.Calculated,
		CalculatedAt: func() *time.Time {
			if row.CalculatedAt.Valid {
				return &row.CalculatedAt.Time
			}
			return nil
		}(),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		DeletedAt: func() *time.Time {
			if row.DeletedAt.Valid {
				return &row.DeletedAt.Time
			}
			return nil
		}(),
	}
}

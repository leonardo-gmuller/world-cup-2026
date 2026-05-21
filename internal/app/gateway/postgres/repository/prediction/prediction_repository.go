package prediction_repository

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/uow"
)

type PredictionRepository struct {
	*sqlc.Queries
}

func NewPredictionRepository(db uow.DBTX) *PredictionRepository {
	return &PredictionRepository{
		Queries: sqlc.New(db),
	}
}

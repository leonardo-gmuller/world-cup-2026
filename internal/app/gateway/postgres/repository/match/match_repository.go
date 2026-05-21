package match_repository

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/uow"
)

type MatchRepository struct {
	*sqlc.Queries
}

func NewMatchRepository(db uow.DBTX) *MatchRepository {
	return &MatchRepository{
		Queries: sqlc.New(db),
	}
}

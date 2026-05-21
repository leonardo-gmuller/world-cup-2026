package group_repository

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/uow"
)

type GroupRepository struct {
	*sqlc.Queries
}

func NewGroupRepository(db uow.DBTX) *GroupRepository {
	return &GroupRepository{
		Queries: sqlc.New(db),
	}
}

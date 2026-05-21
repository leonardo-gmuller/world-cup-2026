package jobcontrol_repository

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/uow"
)

type JobControlRepository struct {
	*sqlc.Queries
}

func NewJobControlRepository(db uow.DBTX) *JobControlRepository {
	return &JobControlRepository{
		Queries: sqlc.New(db),
	}
}

package user_repository

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/uow"
)

type UserRepository struct {
	*sqlc.Queries
}

func NewUserRepository(db uow.DBTX) *UserRepository {
	return &UserRepository{
		Queries: sqlc.New(db),
	}
}

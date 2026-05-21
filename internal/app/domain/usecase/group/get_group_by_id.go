package group_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *GroupUseCase) GetGroupByID(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	return u.repo.GetGroupByID(ctx, id)
}

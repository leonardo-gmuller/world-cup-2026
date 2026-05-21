package group_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *GroupUseCase) ListGroupsByUserID(ctx context.Context, userID int64) ([]entity.Group, error) {
	return u.repo.ListGroupsByUserID(ctx, userID)
}

package group_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *GroupUseCase) GetGroupByInviteCode(ctx context.Context, inviteCode string) (*entity.Group, error) {
	return u.repo.GetGroupByInviteCode(ctx, inviteCode)
}

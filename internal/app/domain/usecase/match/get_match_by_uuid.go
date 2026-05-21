package match_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) GetMatchByUUID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Match, error) {
	return u.repo.GetMatchByUUID(ctx, id)
}

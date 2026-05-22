package match_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) GetMatchByUUID(
	ctx context.Context,
	id string,
) (*entity.Match, error) {
	matchUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return u.repo.GetMatchByUUID(ctx, matchUUID)
}

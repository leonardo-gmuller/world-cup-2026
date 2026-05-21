package match_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) GetMatchByID(
	ctx context.Context,
	id int64,
) (*entity.Match, error) {
	return u.repo.GetMatchByID(ctx, id)
}

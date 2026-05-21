package match_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) ListMatches(ctx context.Context) ([]entity.Match, error) {
	return u.repo.ListMatches(ctx)
}

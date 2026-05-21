package match_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) ListMatchesByStage(
	ctx context.Context,
	stage string,
) ([]entity.Match, error) {
	return u.repo.ListMatchesByStage(ctx, stage)
}

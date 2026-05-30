package match_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) CalculateMatches(ctx context.Context) ([]entity.Match, error) {
	matches, err := u.repo.ListFinishedOrLiveMatches(ctx)
	if err != nil {
		return nil, err
	}

	return matches, nil
}

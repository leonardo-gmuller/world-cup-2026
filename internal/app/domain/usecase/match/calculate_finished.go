package match_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) CalculateFinishedMatches(ctx context.Context) ([]entity.Match, error) {
	matches, err := u.repo.ListFinishedMatchesToCalculate(ctx)
	if err != nil {
		return nil, err
	}

	return matches, nil
}

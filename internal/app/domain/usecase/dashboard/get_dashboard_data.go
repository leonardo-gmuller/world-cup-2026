package dashboard_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type HomeDashboardOutput struct {
	NextMatch entity.Match `json:"next_match"`

	PredictionsCount int64 `json:"predictions_count"`
	GroupsCount      int64 `json:"groups_count"`
	MatchesCount     int64 `json:"matches_count"`
}

func (u *DashboardUseCase) GetDashboardData(
	ctx context.Context,
	userID int64,
) (*HomeDashboardOutput, error) {
	nextMatch, err := u.matchRepo.GetNextMatch(ctx)
	if err != nil {
		return nil, err
	}

	predictionsCount, err := u.predictionRepo.CountPredictionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	groupsCount, err := u.groupRepo.CountGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	matchesCount, err := u.matchRepo.CountMatches(ctx)
	if err != nil {
		return nil, err
	}

	return &HomeDashboardOutput{
		NextMatch:        *nextMatch,
		PredictionsCount: predictionsCount,
		GroupsCount:      groupsCount,
		MatchesCount:     matchesCount,
	}, nil
}

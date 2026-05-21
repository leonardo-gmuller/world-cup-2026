package dashboard_usecase

import "context"

type DashboardUseCase struct {
	groupRepo      groupRepository
	matchRepo      matchRepository
	predictionRepo predictionRepository
	// rankingRepo    rankingRepository
}

type DashboardUseCaseInterface interface {
	GetDashboardData(
		ctx context.Context,
		userID int64,
	) (*HomeDashboardOutput, error)
}

func NewDashboardUseCase(
	groupRepo groupRepository,
	matchRepo matchRepository,
	predictionRepo predictionRepository,
	// rankingRepo rankingRepository,
) *DashboardUseCase {
	return &DashboardUseCase{
		groupRepo:      groupRepo,
		matchRepo:      matchRepo,
		predictionRepo: predictionRepo,
		// rankingRepo:    rankingRepo,
	}
}

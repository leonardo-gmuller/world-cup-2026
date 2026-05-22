package dashboard_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type groupRepository interface {
	CountGroupsByUserID(
		ctx context.Context,
		userID int64,
	) (int64, error)
}

type matchRepository interface {
	GetNextMatch(
		ctx context.Context,
	) (*entity.Match, error)
	CountMatches(
		ctx context.Context,
	) (int64, error)
}

type predictionRepository interface {
	CountPredictionsByUserID(
		ctx context.Context,
		userID int64,
	) (int64, error)

	GetBestRankingByUserID(
		ctx context.Context,
		userID int64,
	) (*entity.UserRanking, error)
}

// type rankingRepository interface {
// 	GetUserRanking(
// 		ctx context.Context,
// 		userID int64,
// 	) (*entity.UserRanking, error)
// }

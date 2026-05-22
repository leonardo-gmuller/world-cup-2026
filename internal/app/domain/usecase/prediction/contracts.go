package prediction_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type predictionRepository interface {
	UpsertPrediction(
		ctx context.Context,
		prediction entity.Prediction,
	) (*entity.Prediction, error)

	GetPrediction(
		ctx context.Context,
		groupID int64,
		userID int64,
		matchID int64,
	) (*entity.Prediction, error)

	GetPredictionByID(
		ctx context.Context,
		id uuid.UUID,
	) (*entity.Prediction, error)

	ListPredictionsByUserAndGroup(
		ctx context.Context,
		groupID int64,
		userID int64,
	) ([]entity.Prediction, error)

	ListPredictionsByMatch(
		ctx context.Context,
		matchID int64,
	) ([]entity.Prediction, error)

	UpdatePredictionPoints(
		ctx context.Context,
		predictionID int64,
		points float64,
	) error

	GetStageWeight(
		ctx context.Context,
		stage string,
	) (*entity.StageWeight, error)

	GetGroupRanking(
		ctx context.Context,
		groupID int64,
	) ([]RankingItemOutput, error)
}

type matchRepository interface {
	GetMatchByID(
		ctx context.Context,
		id int64,
	) (*entity.Match, error)

	GetMatchByUUID(
		ctx context.Context,
		id uuid.UUID,
	) (*entity.Match, error)

	ListFinishedMatches(ctx context.Context) ([]entity.Match, error)
}

type groupRepository interface {
	GetGroupByID(
		ctx context.Context,
		id uuid.UUID,
	) (*entity.Group, error)
}

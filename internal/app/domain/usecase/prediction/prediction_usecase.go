package prediction_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type PredictionUseCase struct {
	predictionRepo  predictionRepository
	matchRepository matchRepository
}

type PredictionUseCaseInterface interface {
	SavePrediction(ctx context.Context, in SavePredictionInput) (*entity.Prediction, error)
	ListPredictionsByUserAndGroup(
		ctx context.Context,
		groupID int64,
		userID int64,
	) ([]entity.Prediction, error)

	CalculateMatchPredictions(
		ctx context.Context,
		matchID int64,
	) error

	GetGroupRanking(
		ctx context.Context,
		groupID int64,
	) ([]RankingItemOutput, error)

	GetPredictionByMatchAndUser(
		ctx context.Context,
		groupID int64,
		userID int64,
		matchID int64,
	) (*entity.Prediction, error)

	GetPredictionByID(ctx context.Context, id uuid.UUID) (*entity.Prediction, error)
}

func NewPredictionUseCase(
	repo predictionRepository,
	matchRepo matchRepository,
) PredictionUseCaseInterface {
	return &PredictionUseCase{
		predictionRepo:  repo,
		matchRepository: matchRepo,
	}
}

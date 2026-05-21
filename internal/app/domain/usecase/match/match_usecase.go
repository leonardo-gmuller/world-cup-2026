package match_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type MatchUseCase struct {
	repo   matchRepository
	client footballAPIClient
}

type MatchUseCaseInterface interface {
	ListMatches(ctx context.Context) ([]entity.Match, error)
	ListMatchesByStage(ctx context.Context, stage string) ([]entity.Match, error)
	GetMatchByID(ctx context.Context, id int64) (*entity.Match, error)
	GetMatchByUUID(ctx context.Context, id uuid.UUID) (*entity.Match, error)
	ImportMatches(ctx context.Context) error
	CalculateFinishedMatches(ctx context.Context) ([]entity.Match, error)
}

func NewMatchUseCase(
	repo matchRepository,
	client footballAPIClient,
) MatchUseCaseInterface {
	return &MatchUseCase{
		repo:   repo,
		client: client,
	}
}

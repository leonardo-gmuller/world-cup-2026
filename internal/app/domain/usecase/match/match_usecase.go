package match_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type MatchUseCase struct {
	repo       matchRepository
	client     footballAPIClient
	liveClient liveScoreClient
}

type MatchUseCaseInterface interface {
	ListMatches(ctx context.Context) ([]entity.Match, error)
	ListMatchesByStage(ctx context.Context, stage string) ([]entity.Match, error)
	GetMatchByID(ctx context.Context, id int64) (*entity.Match, error)
	GetMatchByUUID(ctx context.Context, id string) (*entity.Match, error)
	ImportMatches(ctx context.Context) error
	CalculateMatches(ctx context.Context) ([]entity.Match, error)
	HasLiveMatches(ctx context.Context) (bool, error)
	SyncLiveResults(ctx context.Context) error
}

func NewMatchUseCase(
	repo matchRepository,
	client footballAPIClient,
	liveClient liveScoreClient,
) MatchUseCaseInterface {
	return &MatchUseCase{
		repo:       repo,
		client:     client,
		liveClient: liveClient,
	}
}

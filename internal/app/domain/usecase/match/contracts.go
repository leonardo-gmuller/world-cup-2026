package match_usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type ExternalTeamOutput struct {
	ExternalID string
	Name       string
	ShortName  string
	Code       string
	FlagURL    string
}

type ExternalMatchOutput struct {
	ExternalID string
	Stage      string
	GroupName  *string
	HomeTeam   *ExternalTeamOutput
	AwayTeam   *ExternalTeamOutput
	StartsAt   time.Time
	HomeScore  *int
	AwayScore  *int
	Status     string
}

type matchRepository interface {
	ListMatches(ctx context.Context) ([]entity.Match, error)
	ListMatchesByStage(ctx context.Context, stage string) ([]entity.Match, error)
	GetMatchByID(ctx context.Context, id int64) (*entity.Match, error)
	GetMatchByUUID(ctx context.Context, id uuid.UUID) (*entity.Match, error)

	UpsertTeam(ctx context.Context, team entity.Team) (*entity.Team, error)
	UpsertMatch(ctx context.Context, match entity.Match) (*entity.Match, error)

	ListFinishedOrLiveMatches(ctx context.Context) ([]entity.Match, error)

	HasLiveMatches(ctx context.Context) (bool, error)

	FindMatchForLiveSync(
		ctx context.Context,
		startsAt time.Time,
		homeTeamName string,
		awayTeamName string,
	) (*entity.Match, error)

	UpdateLiveResult(
		ctx context.Context,
		matchID int64,
		apiFootballID int64,
		homeScore *int,
		awayScore *int,
		status string,
	) (*entity.Match, error)

	GetMatchByExternalID(
		ctx context.Context,
		externalID string,
	) (*entity.Match, error)

	HasMatchesToSyncLiveResults(
		ctx context.Context,
	) (bool, error)
}

type footballAPIClient interface {
	FetchWorldCupMatches(ctx context.Context) ([]ExternalMatchOutput, error)
}

type liveScoreClient interface {
	FetchTodayMatches(
		ctx context.Context,
		date time.Time,
	) ([]ExternalMatchOutput, error)
}

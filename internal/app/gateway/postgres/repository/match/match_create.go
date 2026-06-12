package match_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/util"
)

func (r *MatchRepository) UpsertTeam(
	ctx context.Context,
	team entity.Team,
) (*entity.Team, error) {
	row, err := r.Queries.UpsertTeam(ctx, sqlc.UpsertTeamParams{
		Uuid:       team.UUID,
		ExternalID: util.TextPtr(team.ExternalID),
		Name:       team.Name,
		ShortName:  util.TextPtr(team.ShortName),
		Code:       util.TextPtr(team.Code),
		FlagUrl:    util.TextPtr(team.FlagURL),
	})
	if err != nil {
		return nil, err
	}

	return mapTeam(row), nil
}

func (r *MatchRepository) UpsertMatch(
	ctx context.Context,
	match entity.Match,
) (*entity.Match, error) {
	row, err := r.Queries.UpsertMatch(ctx, sqlc.UpsertMatchParams{
		Uuid:         match.UUID,
		ExternalID:   util.TextPtr(match.ExternalID),
		Stage:        match.Stage,
		GroupName:    util.TextPtr(match.GroupName),
		HomeTeamID:   util.Int8Ptr(match.HomeTeamID),
		AwayTeamID:   util.Int8Ptr(match.AwayTeamID),
		HomeTeamName: util.TextPtr(match.HomeTeamName),
		AwayTeamName: util.TextPtr(match.AwayTeamName),
		StartsAt:     pgtype.Timestamptz{Time: match.StartsAt, Valid: true},
		HomeScore:    util.Int4Ptr(match.HomeScore),
		AwayScore:    util.Int4Ptr(match.AwayScore),
		Status:       match.Status,
		WinnerTeamID: util.Int8Ptr(match.WinnerTeamID),
	})
	if err != nil {
		return nil, err
	}

	return mapMatch(row), nil
}

func (r *MatchRepository) UpdateLiveResult(
	ctx context.Context,
	matchID int64,
	apiFootballID int64,
	homeScore *int,
	awayScore *int,
	status string,
) (*entity.Match, error) {
	match, err := r.Queries.UpdateLiveResult(ctx, sqlc.UpdateLiveResultParams{
		ID:            matchID,
		ApiFootballID: pgtype.Int8{Int64: apiFootballID, Valid: true},
		HomeScore:     pgtype.Int4{Int32: int32(*homeScore), Valid: homeScore != nil},
		AwayScore:     pgtype.Int4{Int32: int32(*awayScore), Valid: awayScore != nil},
		Status:        status,
	})
	if err != nil {
		return nil, err
	}

	return mapMatch(match), nil
}

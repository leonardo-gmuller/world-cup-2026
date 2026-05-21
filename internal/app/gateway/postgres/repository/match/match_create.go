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

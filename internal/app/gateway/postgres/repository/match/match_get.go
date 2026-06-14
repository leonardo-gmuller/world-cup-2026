package match_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func (r *MatchRepository) GetMatchByID(
	ctx context.Context,
	id int64,
) (*entity.Match, error) {
	row, err := r.Queries.GetMatchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapMatchByIDRow(row), nil
}

func (r *MatchRepository) GetMatchByUUID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Match, error) {
	row, err := r.Queries.GetMatchByUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapMatchByUUIDRow(row), nil
}

func (r *MatchRepository) ListMatches(
	ctx context.Context,
) ([]entity.Match, error) {
	rows, err := r.Queries.ListMatches(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]entity.Match, 0, len(rows))
	for _, row := range rows {
		items = append(items, *mapListMatchesRow(row))
	}

	return items, nil
}

func (r *MatchRepository) ListMatchesByStage(
	ctx context.Context,
	stage string,
) ([]entity.Match, error) {
	rows, err := r.Queries.ListMatchesByStage(ctx, stage)
	if err != nil {
		return nil, err
	}

	items := make([]entity.Match, 0, len(rows))
	for _, row := range rows {
		items = append(items, *mapMatch(row))
	}

	return items, nil
}

func (r *MatchRepository) CountMatches(
	ctx context.Context,
) (int64, error) {

	total, err := r.Queries.CountMatches(ctx)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *MatchRepository) GetNextMatch(
	ctx context.Context,
) (*entity.Match, error) {

	row, err := r.Queries.GetNextMatch(ctx)
	if err != nil {
		return nil, err
	}

	match := mapGetNextMatchRow(row)

	return match, nil
}

func (r *MatchRepository) ListFinishedOrLiveMatches(ctx context.Context) ([]entity.Match, error) {
	rows, err := r.Queries.ListFinishedOrLiveMatchesToCalculate(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]entity.Match, 0, len(rows))
	for _, row := range rows {
		items = append(items, *mapMatch(row))
	}

	return items, nil
}

func (r *MatchRepository) HasLiveMatches(ctx context.Context) (bool, error) {
	row, err := r.Queries.HasLiveMatches(ctx)
	if err != nil {
		return false, err
	}

	return row, nil
}

func (r *MatchRepository) FindMatchForLiveSync(
	ctx context.Context,
	startsAt time.Time,
	homeTeamName string,
	awayTeamName string,
) (*entity.Match, error) {
	match, err := r.Queries.FindMatchForLiveSync(ctx, sqlc.FindMatchForLiveSyncParams{
		StartsAt:     pgtype.Timestamptz{Time: startsAt, Valid: true},
		HomeTeamName: homeTeamName,
		AwayTeamName: awayTeamName,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return mapMatch(match), nil
}

func (r *MatchRepository) GetMatchByExternalID(
	ctx context.Context,
	externalID string,
) (*entity.Match, error) {
	match, err := r.Queries.GetMatchByExternalID(ctx, pgtype.Text{
		String: externalID,
		Valid:  true,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	entityMatch := mapMatch(match)

	return entityMatch, nil
}

func (r *MatchRepository) ListMatchesToSyncLiveResults(
	ctx context.Context,
) ([]entity.Match, error) {
	rows, err := r.Queries.ListMatchesToSyncLiveResults(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]entity.Match, 0, len(rows))
	for _, row := range rows {
		items = append(items, *mapMatch(row))
	}

	return items, nil
}

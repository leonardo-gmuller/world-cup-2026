package match_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
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

func (r *MatchRepository) ListFinishedMatchesToCalculate(
	ctx context.Context,
) ([]entity.Match, error) {
	rows, err := r.Queries.ListFinishedMatchesToCalculate(ctx)
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

func (r *MatchRepository) ListFinishedMatches(ctx context.Context) ([]entity.Match, error) {
	rows, err := r.Queries.ListFinishedMatchesToCalculate(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]entity.Match, 0, len(rows))
	for _, row := range rows {
		items = append(items, *mapMatch(row))
	}

	return items, nil
}

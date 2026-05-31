package prediction_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	prediction_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/prediction"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func (r *PredictionRepository) GetPrediction(
	ctx context.Context,
	groupID int64,
	userID int64,
	matchID int64,
) (*entity.Prediction, error) {
	row, err := r.Queries.GetPrediction(ctx, sqlc.GetPredictionParams{
		GroupID: groupID,
		UserID:  userID,
		MatchID: matchID,
	})
	if err != nil {
		return nil, err
	}

	return mapPrediction(row), nil
}

func (r *PredictionRepository) GetPredictionByID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Prediction, error) {
	row, err := r.Queries.GetPredictionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapPrediction(row), nil
}

func (r *PredictionRepository) ListPredictionsByUserAndGroup(
	ctx context.Context,
	groupID int64,
	userID int64,
) ([]entity.Prediction, error) {
	rows, err := r.Queries.ListPredictionsByUserAndGroup(ctx, sqlc.ListPredictionsByUserAndGroupParams{
		GroupID: groupID,
		UserID:  userID,
	})
	if err != nil {
		return nil, err
	}

	items := make([]entity.Prediction, 0, len(rows))
	for _, row := range rows {
		items = append(items, *mapPredictionsByGroup(row))
	}

	return items, nil
}

func (r *PredictionRepository) ListPredictionsByMatch(
	ctx context.Context,
	matchID int64,
) ([]entity.Prediction, error) {
	rows, err := r.Queries.ListPredictionsByMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	items := make([]entity.Prediction, 0, len(rows))
	for _, row := range rows {
		items = append(items, *mapPrediction(row))
	}

	return items, nil
}

func (r *PredictionRepository) GetGroupRanking(
	ctx context.Context,
	groupID int64,
) ([]prediction_usecase.RankingItemOutput, error) {
	rows, err := r.Queries.GetGroupRanking(ctx, groupID)
	if err != nil {
		return nil, err
	}

	items := make([]prediction_usecase.RankingItemOutput, 0, len(rows))
	for _, row := range rows {
		items = append(items, prediction_usecase.RankingItemOutput{
			UserID:           row.UserID,
			UserUUID:         row.UserUuid,
			Name:             row.Name,
			TotalPoints:      row.TotalPoints,
			PredictionsCount: row.PredictionsCount,
		})
	}

	return items, nil
}

func (r *PredictionRepository) GetStageWeight(
	ctx context.Context,
	stage string,
) (*entity.StageWeight, error) {
	row, err := r.Queries.GetStageWeight(ctx, stage)
	if err != nil {
		return nil, err
	}

	return &entity.StageWeight{
		Stage:      row.Stage,
		Weight:     row.Weight,
		OrderIndex: int(row.OrderIndex),
		CreatedAt:  row.CreatedAt.Time,
	}, nil
}

func (r *PredictionRepository) CountPredictionsByUserID(
	ctx context.Context,
	userID int64,
) (int64, error) {

	total, err := r.Queries.CountPredictionsByUserID(ctx, userID)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PredictionRepository) GetBestRankingByUserID(
	ctx context.Context,
	userID int64,
) (*entity.UserRanking, error) {
	groupIDs, err := r.Queries.ListGroupIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var best *entity.UserRanking

	for _, groupID := range groupIDs {
		rows, err := r.Queries.GetGroupRanking(ctx, groupID)
		if err != nil {
			return nil, err
		}

		totalPlayers := len(rows)

		for index, row := range rows {
			if row.UserID != userID {
				continue
			}

			position := index + 1

			if best == nil ||
				position < best.Position ||
				(position == best.Position && int64(row.TotalPoints) > best.TotalPoints) {
				best = &entity.UserRanking{
					GroupID:      groupID,
					UserID:       userID,
					Position:     position,
					TotalPlayers: int64(totalPlayers),
					TotalPoints:  int64(row.TotalPoints),
				}
			}
		}
	}

	return best, nil
}

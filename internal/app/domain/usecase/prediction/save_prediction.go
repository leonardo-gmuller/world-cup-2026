package prediction_usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/erring"
)

type SavePredictionInput struct {
	GroupID   string
	UserID    int64
	MatchID   string
	HomeScore int
	AwayScore int
}

func (u *PredictionUseCase) SavePrediction(
	ctx context.Context,
	in SavePredictionInput,
) (*entity.Prediction, error) {

	matchUUID, err := uuid.Parse(in.MatchID)
	if err != nil {
		return nil, err
	}

	match, err := u.matchRepository.GetMatchByUUID(ctx, matchUUID)
	if err != nil {
		return nil, err
	}

	groupUUID, err := uuid.Parse(in.GroupID)
	if err != nil {
		return nil, err
	}

	group, err := u.groupRepository.GetGroupByID(ctx, groupUUID)
	if err != nil {
		return nil, err
	}

	limitTime := match.StartsAt.Add(-5 * time.Minute)

	if time.Now().After(limitTime) {
		return nil, erring.ErrPredictionClosed
	}

	prediction := entity.Prediction{
		UUID:      uuid.New(),
		GroupID:   group.ID,
		UserID:    in.UserID,
		MatchID:   match.ID,
		HomeScore: in.HomeScore,
		AwayScore: in.AwayScore,
	}

	newPrediction, err := u.predictionRepo.UpsertPrediction(ctx, prediction)
	if err != nil {
		return nil, err
	}

	newPrediction.GroupUUID = group.UUID
	newPrediction.MatchUUID = match.UUID

	return newPrediction, nil
}

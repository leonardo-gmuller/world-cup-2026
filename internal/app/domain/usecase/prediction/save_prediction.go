package prediction_usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/erring"
)

type SavePredictionInput struct {
	GroupID   int64
	UserID    int64
	MatchID   int64
	HomeScore int
	AwayScore int
}

func (u *PredictionUseCase) SavePrediction(
	ctx context.Context,
	in SavePredictionInput,
) (*entity.Prediction, error) {

	match, err := u.matchRepository.GetMatchByID(ctx, in.MatchID)
	if err != nil {
		return nil, err
	}

	limitTime := match.StartsAt.Add(-5 * time.Minute)

	if time.Now().After(limitTime) {
		return nil, erring.ErrPredictionClosed
	}

	prediction := entity.Prediction{
		UUID:      uuid.New(),
		GroupID:   in.GroupID,
		UserID:    in.UserID,
		MatchID:   in.MatchID,
		HomeScore: in.HomeScore,
		AwayScore: in.AwayScore,
	}

	return u.predictionRepo.UpsertPrediction(ctx, prediction)
}

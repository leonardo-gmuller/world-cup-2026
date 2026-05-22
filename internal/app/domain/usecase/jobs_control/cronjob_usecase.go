package jobcontrol_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
	match_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/match"
	prediction_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/prediction"
)

type JobControlUsecaseInterface interface {
	CreateJobsControl(ctx context.Context, jobID types.Job) error
	UpdateJobsControl(ctx context.Context, jobID types.Job) error
	ImportMatches(ctx context.Context) error
	CalculateMatchPredictions(ctx context.Context) error
}

type JobControlUsecase struct {
	jobsControlRepository jobsControlRepository
	matchUsecase          match_usecase.MatchUseCaseInterface
	predictionUsecase     prediction_usecase.PredictionUseCaseInterface
}

func NewJobControlUsecase(
	repo jobsControlRepository,
	matchUsecase match_usecase.MatchUseCaseInterface,
	predictionUsecase prediction_usecase.PredictionUseCaseInterface,
) JobControlUsecaseInterface {
	return &JobControlUsecase{
		jobsControlRepository: repo,
		matchUsecase:          matchUsecase,
		predictionUsecase:     predictionUsecase,
	}
}

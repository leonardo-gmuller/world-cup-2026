package jobcontrol_repository

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

func (r *JobControlRepository) GetByJob(ctx context.Context, job types.Job) (*entity.JobControl, error) {
	jobControl, err := r.Queries.GetJobControl(ctx, string(job))
	if err != nil {
		return nil, err
	}

	return &entity.JobControl{
		Job:            types.Job(jobControl.Job),
		LastSuccessRun: &jobControl.LastSuccessRun.Time,
	}, nil
}

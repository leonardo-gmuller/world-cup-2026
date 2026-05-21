package jobcontrol_repository

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

func (r *JobControlRepository) Create(ctx context.Context, job types.Job) error {
	err := r.Queries.CreateJobControl(ctx, string(job))
	if err != nil {
		return err
	}

	return nil
}

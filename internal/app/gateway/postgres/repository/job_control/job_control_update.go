package jobcontrol_repository

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

func (r *JobControlRepository) Update(ctx context.Context, job types.Job) error {
	err := r.Queries.UpdateJobControl(ctx, string(job))
	if err != nil {
		return err
	}

	return nil
}

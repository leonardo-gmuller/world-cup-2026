package jobcontrol_usecase

import (
	"context"
	"fmt"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

func (u *JobControlUsecase) CreateJobsControl(ctx context.Context, jobID types.Job) error {
	const operation = "UseCase.CreateJobsControl"

	err := u.jobsControlRepository.Create(ctx, jobID)
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}

	return nil
}

func (u *JobControlUsecase) UpdateJobsControl(ctx context.Context, jobID types.Job) error {
	const operation = "UseCase.UpdateJobsControl"

	err := u.jobsControlRepository.Update(ctx, jobID)
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}

	return nil
}

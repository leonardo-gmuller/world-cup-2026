package jobcontrol_usecase

import (
	"context"
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

func (u *JobControlUsecase) ShouldImportMatches(ctx context.Context, duration time.Duration) (bool, error) {
	job, err := u.jobsControlRepository.GetByJob(ctx, types.ImportMatches)
	if err != nil {
		return false, err
	}

	if job == nil {
		return true, nil
	}

	if job.LastSuccessRun == nil {
		return true, nil
	}

	return time.Since(*job.LastSuccessRun) > duration, nil
}

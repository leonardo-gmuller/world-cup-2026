package jobcontrol_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

type jobsControlRepository interface {
	Create(ctx context.Context, job types.Job) error
	Update(ctx context.Context, job types.Job) error
	GetByJob(ctx context.Context, job types.Job) (*entity.JobControl, error)
}

package entity

import (
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/types"
)

type JobControl struct {
	Job            types.Job
	LastSuccessRun *time.Time
}

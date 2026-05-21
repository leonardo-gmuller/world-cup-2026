package entity

import (
	"time"

	"github.com/google/uuid"
)

type Prediction struct {
	ID           int64
	UUID         uuid.UUID
	GroupID      int64
	UserID       int64
	MatchID      int64
	HomeScore    int
	AwayScore    int
	Points       float64
	Calculated   bool
	CalculatedAt *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

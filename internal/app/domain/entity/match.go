package entity

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID              int64
	UUID            uuid.UUID
	ExternalID      *string
	Stage           string
	GroupName       *string
	HomeTeamID      *int64
	AwayTeamID      *int64
	HomeTeamName    *string
	AwayTeamName    *string
	HomeTeamFlagURL *string
	AwayTeamFlagURL *string
	StartsAt        time.Time
	HomeScore       *int
	AwayScore       *int
	Status          string
	WinnerTeamID    *int64
	ImportedAt      *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

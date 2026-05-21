package entity

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID         int64
	UUID       uuid.UUID
	ExternalID *string
	Name       string
	ShortName  *string
	Code       *string
	FlagURL    *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           int64
	UUID         uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

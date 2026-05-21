package entity

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID          int64
	UUID        uuid.UUID
	Name        string
	Description *string
	OwnerID     int64
	InviteCode  string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type GroupMember struct {
	ID        int64
	UUID      uuid.UUID
	GroupID   int64
	UserID    int64
	Role      string
	JoinedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

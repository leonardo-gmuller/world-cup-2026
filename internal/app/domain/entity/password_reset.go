package entity

import "time"

type PasswordResetToken struct {
	ID        int64
	UserID    int64
	Token     string
	Used      bool
	ExpiresAt time.Time
	CreatedAt time.Time
	UsedAt    *time.Time
}

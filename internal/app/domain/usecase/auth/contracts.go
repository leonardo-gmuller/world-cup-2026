package auth_usecase

import (
	"context"
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type authRepository interface {
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUserPassword(ctx context.Context, userID int64, password string) error
}

type passwordResetRepository interface {
	CreatePasswordResetToken(
		ctx context.Context,
		userID int64,
		token string,
		expiresAt time.Time,
	) (*entity.PasswordResetToken, error)

	GetPasswordResetTokenByToken(
		ctx context.Context,
		token string,
	) (*entity.PasswordResetToken, error)

	InvalidatePasswordResetTokensByUserID(
		ctx context.Context,
		userID int64,
	) error

	UsePasswordResetToken(
		ctx context.Context,
		id int64,
	) error
}

type hashService interface {
	Hash(password string) (string, error)
	Compare(hash string, password string) bool
}

type jwtService interface {
	Generate(user entity.User) (string, error)
}

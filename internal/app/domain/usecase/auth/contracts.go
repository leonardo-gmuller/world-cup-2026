package auth_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type authRepository interface {
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type hashService interface {
	Hash(password string) (string, error)
	Compare(hash string, password string) bool
}

type jwtService interface {
	Generate(user entity.User) (string, error)
}

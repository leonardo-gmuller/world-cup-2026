package user_repository

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user entity.User,
) (*entity.User, error) {
	row, err := r.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		Uuid:         user.UUID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {
		return nil, err
	}

	return mapUser(row), nil
}

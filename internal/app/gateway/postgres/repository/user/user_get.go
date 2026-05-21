package user_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (r *UserRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	row, err := r.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return mapUser(row), nil
}

func (r *UserRepository) GetUserByUUID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.User, error) {
	row, err := r.Queries.GetUserByUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapUser(row), nil
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int64,
) (*entity.User, error) {
	row, err := r.Queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapUser(row), nil
}

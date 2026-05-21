package user_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *UserUseCase) GetUserByID(
	ctx context.Context,
	id int64,
) (*entity.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u *UserUseCase) GetUserByUUID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.User, error) {
	return u.repo.GetUserByUUID(ctx, id)
}

func (u *UserUseCase) GetUserByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

package auth_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/erring"
)

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

type AuthOutput struct {
	Token string
	User  entity.User
}

func (u *AuthUseCase) Register(
	ctx context.Context,
	in RegisterInput,
) (*AuthOutput, error) {
	existingUser, _ := u.repo.GetUserByEmail(ctx, in.Email)
	if existingUser != nil {
		return nil, erring.ErrEmailAlreadyUsed
	}

	passwordHash, err := u.hash.Hash(in.Password)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		UUID:         uuid.New(),
		Name:         in.Name,
		Email:        in.Email,
		PasswordHash: passwordHash,
		IsActive:     true,
	}

	createdUser, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := u.jwt.Generate(*createdUser)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: token,
		User:  *createdUser,
	}, nil
}

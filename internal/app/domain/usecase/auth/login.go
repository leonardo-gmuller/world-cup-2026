package auth_usecase

import (
	"context"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/erring"
)

type LoginInput struct {
	Email    string
	Password string
}

func (u *AuthUseCase) Login(
	ctx context.Context,
	in LoginInput,
) (*AuthOutput, error) {
	user, err := u.repo.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return nil, erring.ErrInvalidCredentials
	}

	if user == nil || !user.IsActive {
		return nil, erring.ErrInvalidCredentials
	}

	if !u.hash.Compare(user.PasswordHash, in.Password) {
		return nil, erring.ErrInvalidCredentials
	}

	token, err := u.jwt.Generate(*user)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: token,
		User:  *user,
	}, nil
}

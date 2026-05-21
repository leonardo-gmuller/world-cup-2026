package auth_usecase

import "context"

type AuthUseCase struct {
	repo authRepository
	hash hashService
	jwt  jwtService
}

type AuthUseCaseInterface interface {
	Register(ctx context.Context, in RegisterInput) (*AuthOutput, error)
	Login(ctx context.Context, in LoginInput) (*AuthOutput, error)
}

func NewAuthUseCase(
	repo authRepository,
	hash hashService,
	jwt jwtService,
) AuthUseCaseInterface {
	return &AuthUseCase{
		repo: repo,
		hash: hash,
		jwt:  jwt,
	}
}

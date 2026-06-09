package auth_usecase

import "context"

type AuthUseCase struct {
	repo              authRepository
	passwordResetRepo passwordResetRepository
	hash              hashService
	jwt               jwtService
}

type AuthUseCaseInterface interface {
	Register(ctx context.Context, in RegisterInput) (*AuthOutput, error)
	Login(ctx context.Context, in LoginInput) (*AuthOutput, error)
	ForgotPassword(ctx context.Context, in ForgotPasswordInput) (*ForgotPasswordOutput, error)
	ResetPassword(ctx context.Context, in ResetPasswordInput) error
}

func NewAuthUseCase(
	repo authRepository,
	passwordResetRepo passwordResetRepository,
	hash hashService,
	jwt jwtService,
) AuthUseCaseInterface {
	return &AuthUseCase{
		repo:              repo,
		passwordResetRepo: passwordResetRepo,
		hash:              hash,
		jwt:               jwt,
	}
}

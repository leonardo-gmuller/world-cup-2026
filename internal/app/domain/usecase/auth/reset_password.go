package auth_usecase

import (
	"context"
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/erring"
)

type ResetPasswordInput struct {
	Token                string
	Password             string
	PasswordConfirmation string
}

func (u *AuthUseCase) ResetPassword(
	ctx context.Context,
	in ResetPasswordInput,
) error {
	if in.Password != in.PasswordConfirmation {
		return erring.ErrPasswordsDontMatch
	}

	resetToken, err := u.passwordResetRepo.GetPasswordResetTokenByToken(ctx, in.Token)
	if err != nil {
		return erring.ErrInvalidResetToken
	}

	if resetToken.Used {
		return erring.ErrUsedResetToken
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return erring.ErrExpiredResetToken
	}

	hashedPassword, err := u.hash.Hash(in.Password)
	if err != nil {
		return err
	}

	if err := u.repo.UpdateUserPassword(ctx, resetToken.UserID, hashedPassword); err != nil {
		return err
	}

	if err := u.passwordResetRepo.UsePasswordResetToken(ctx, resetToken.ID); err != nil {
		return err
	}

	return nil
}

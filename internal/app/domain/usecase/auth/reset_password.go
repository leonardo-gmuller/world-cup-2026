package auth_usecase

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidResetToken  = errors.New("token inválido")
	ErrExpiredResetToken  = errors.New("token expirado")
	ErrUsedResetToken     = errors.New("token já utilizado")
	ErrPasswordsDontMatch = errors.New("as senhas não conferem")
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
		return ErrPasswordsDontMatch
	}

	resetToken, err := u.passwordResetRepo.GetPasswordResetTokenByToken(ctx, in.Token)
	if err != nil {
		return ErrInvalidResetToken
	}

	if resetToken.Used {
		return ErrUsedResetToken
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return ErrExpiredResetToken
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

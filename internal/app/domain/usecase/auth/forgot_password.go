package auth_usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type ForgotPasswordInput struct {
	Email string
}

type ForgotPasswordOutput struct {
	Token    string
	ResetURL string
}

func (u *AuthUseCase) ForgotPassword(
	ctx context.Context,
	in ForgotPasswordInput,
) (*ForgotPasswordOutput, error) {
	user, err := u.repo.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	if err := u.passwordResetRepo.InvalidatePasswordResetTokensByUserID(ctx, user.ID); err != nil {
		return nil, err
	}

	token, err := generateResetToken()
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(1 * time.Hour)

	if _, err := u.passwordResetRepo.CreatePasswordResetToken(ctx, user.ID, token, expiresAt); err != nil {
		return nil, err
	}

	return &ForgotPasswordOutput{
		Token:    token,
		ResetURL: fmt.Sprintf("/reset-password?token=%s", token),
	}, nil
}

func generateResetToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

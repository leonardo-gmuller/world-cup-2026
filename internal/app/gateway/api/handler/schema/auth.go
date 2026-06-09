package schema

import (
	"errors"
	"fmt"

	auth_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/auth"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *RegisterRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	if r.Email == "" {
		return fmt.Errorf("email is required")
	}

	if r.Password == "" {
		return fmt.Errorf("password is required")
	}

	if len(r.Password) < 6 {
		return fmt.Errorf("password must have at least 6 characters")
	}

	return nil
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) Validate() error {
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}

	if r.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func AuthResponseFromUseCase(out *auth_usecase.AuthOutput) AuthResponse {
	return AuthResponse{
		Token: out.Token,
		User: UserResponse{
			ID:        out.User.UUID.String(),
			Name:      out.User.Name,
			Email:     out.User.Email,
			CreatedAt: out.User.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

func (r ForgotPasswordRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email é obrigatório")
	}

	return nil
}

type ForgotPasswordResponse struct {
	Token    string `json:"token"`
	ResetURL string `json:"reset_url"`
}

func ForgotPasswordResponseFromUseCase(out *auth_usecase.ForgotPasswordOutput) ForgotPasswordResponse {
	return ForgotPasswordResponse{
		Token:    out.Token,
		ResetURL: out.ResetURL,
	}
}

type ResetPasswordRequest struct {
	Token                string `json:"token"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (r ResetPasswordRequest) Validate() error {
	if r.Token == "" {
		return errors.New("token é obrigatório")
	}

	if r.Password == "" {
		return errors.New("senha é obrigatória")
	}

	if len(r.Password) < 6 {
		return errors.New("senha deve ter no mínimo 6 caracteres")
	}

	if r.Password != r.PasswordConfirmation {
		return errors.New("as senhas não conferem")
	}

	return nil
}

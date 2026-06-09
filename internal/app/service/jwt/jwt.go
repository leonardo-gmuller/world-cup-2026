package jwt

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type Service struct {
	auth *jwtauth.JWTAuth
}

func New(secret string) *Service {
	return &Service{
		auth: jwtauth.New("HS256", []byte(secret), nil),
	}
}

func (s *Service) Generate(user entity.User) (string, error) {
	now := time.Now()

	ttl := 70 * 24 * time.Hour

	_, tokenString, err := s.auth.Encode(map[string]any{
		"user_id": user.ID,
		"uuid":    user.UUID.String(),
		"name":    user.Name,
		"email":   user.Email,
		"exp":     now.Add(ttl).Unix(),
		"iat":     now.Unix(),
	})

	return tokenString, err
}

func (s *Service) Auth() *jwtauth.JWTAuth {
	return s.auth
}

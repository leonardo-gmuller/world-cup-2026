package schema

import "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func UserResponseFromEntity(user *entity.User) UserResponse {
	return UserResponse{
		ID:        user.UUID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

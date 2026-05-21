package user_repository

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func mapUser(row sqlc.User) *entity.User {
	return &entity.User{
		ID:           row.ID,
		UUID:         row.Uuid,
		Name:         row.Name,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
		IsActive:     row.IsActive,
		CreatedAt:    row.CreatedAt.Time,
		UpdatedAt:    row.UpdatedAt.Time,
		DeletedAt:    &row.DeletedAt.Time,
	}
}

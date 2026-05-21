package group_repository

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func mapGroup(row sqlc.Group) *entity.Group {
	return &entity.Group{
		ID:          row.ID,
		UUID:        row.Uuid,
		Name:        row.Name,
		Description: &row.Description.String,
		OwnerID:     row.OwnerID,
		InviteCode:  row.InviteCode,
		IsActive:    row.IsActive,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
		DeletedAt:   &row.DeletedAt.Time,
	}
}

func mapGroupMember(row sqlc.GroupMember) *entity.GroupMember {
	return &entity.GroupMember{
		ID:        row.ID,
		UUID:      row.Uuid,
		GroupID:   row.GroupID,
		UserID:    row.UserID,
		Role:      row.Role,
		JoinedAt:  row.JoinedAt.Time,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		DeletedAt: &row.DeletedAt.Time,
	}
}

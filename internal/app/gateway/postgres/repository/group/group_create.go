package group_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/util"
)

func (r *GroupRepository) CreateGroup(
	ctx context.Context,
	group entity.Group,
) (*entity.Group, error) {
	row, err := r.Queries.CreateGroup(ctx, sqlc.CreateGroupParams{
		Uuid:        group.UUID,
		Name:        group.Name,
		Description: pgtype.Text{String: util.DerefOrEmpty(group.Description), Valid: true},
		OwnerID:     group.OwnerID,
		InviteCode:  group.InviteCode,
	})
	if err != nil {
		return nil, err
	}

	return mapGroup(row), nil
}

func (r *GroupRepository) CreateGroupMember(
	ctx context.Context,
	member entity.GroupMember,
) (*entity.GroupMember, error) {
	row, err := r.Queries.CreateGroupMember(ctx, sqlc.CreateGroupMemberParams{
		Uuid:    member.UUID,
		GroupID: member.GroupID,
		UserID:  member.UserID,
		Role:    member.Role,
	})
	if err != nil {
		return nil, err
	}

	return mapGroupMember(row), nil
}

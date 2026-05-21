package group_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func (r *GroupRepository) GetGroupByID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Group, error) {
	row, err := r.Queries.GetGroupByUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapGroup(row), nil
}

func (r *GroupRepository) GetGroupByInviteCode(
	ctx context.Context,
	inviteCode string,
) (*entity.Group, error) {
	row, err := r.Queries.GetGroupByInviteCode(ctx, inviteCode)
	if err != nil {
		return nil, err
	}

	return mapGroup(row), nil
}

func (r *GroupRepository) GetGroupMember(
	ctx context.Context,
	groupID int64,
	userID int64,
) (*entity.GroupMember, error) {
	row, err := r.Queries.GetGroupMember(ctx, sqlc.GetGroupMemberParams{
		GroupID: groupID,
		UserID:  userID,
	})
	if err != nil {
		return nil, err
	}

	return mapGroupMember(row), nil
}

func (r *GroupRepository) ListGroupsByUserID(
	ctx context.Context,
	userID int64,
) ([]entity.Group, error) {
	rows, err := r.Queries.ListGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	groups := make([]entity.Group, 0, len(rows))
	for _, row := range rows {
		groups = append(groups, *mapGroup(row))
	}

	return groups, nil
}

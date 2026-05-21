package group_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type groupRepository interface {
	CreateGroup(ctx context.Context, group entity.Group) (*entity.Group, error)
	GetGroupByID(ctx context.Context, id uuid.UUID) (*entity.Group, error)
	GetGroupByInviteCode(ctx context.Context, inviteCode string) (*entity.Group, error)
	ListGroupsByUserID(ctx context.Context, userID int64) ([]entity.Group, error)

	CreateGroupMember(ctx context.Context, member entity.GroupMember) (*entity.GroupMember, error)
	GetGroupMember(ctx context.Context, groupID int64, userID int64) (*entity.GroupMember, error)
}

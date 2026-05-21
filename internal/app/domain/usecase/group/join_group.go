package group_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/constants"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type JoinGroupInput struct {
	UserID     int64
	InviteCode string
}

func (u *GroupUseCase) JoinGroup(ctx context.Context, in JoinGroupInput) (*entity.GroupMember, error) {
	group, err := u.repo.GetGroupByInviteCode(ctx, in.InviteCode)
	if err != nil {
		return nil, err
	}

	member, err := u.repo.GetGroupMember(ctx, group.ID, in.UserID)
	if err == nil && member != nil {
		return member, nil
	}

	newMember := entity.GroupMember{
		UUID:    uuid.New(),
		GroupID: group.ID,
		UserID:  in.UserID,
		Role:    constants.GroupMemberRoleMember,
	}

	return u.repo.CreateGroupMember(ctx, newMember)
}

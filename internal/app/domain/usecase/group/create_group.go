package group_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/constants"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type CreateGroupInput struct {
	Name        string
	Description *string
	OwnerID     int64
}

func (u *GroupUseCase) CreateGroup(ctx context.Context, in CreateGroupInput) (*entity.Group, error) {
	group := entity.Group{
		UUID:        uuid.New(),
		Name:        in.Name,
		Description: in.Description,
		OwnerID:     in.OwnerID,
		InviteCode:  uuid.NewString(),
		IsActive:    true,
	}

	createdGroup, err := u.repo.CreateGroup(ctx, group)
	if err != nil {
		return nil, err
	}

	member := entity.GroupMember{
		UUID:    uuid.New(),
		GroupID: createdGroup.ID,
		UserID:  in.OwnerID,
		Role:    constants.GroupMemberRoleOwner,
	}

	_, err = u.repo.CreateGroupMember(ctx, member)
	if err != nil {
		return nil, err
	}

	return createdGroup, nil
}

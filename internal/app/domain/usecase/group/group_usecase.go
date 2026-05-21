package group_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type GroupUseCase struct {
	repo groupRepository
}

type GroupUseCaseInterface interface {
	CreateGroup(ctx context.Context, in CreateGroupInput) (*entity.Group, error)
	GetGroupByID(ctx context.Context, id uuid.UUID) (*entity.Group, error)
	GetGroupByInviteCode(ctx context.Context, inviteCode string) (*entity.Group, error)
	ListGroupsByUserID(ctx context.Context, userID int64) ([]entity.Group, error)
	JoinGroup(ctx context.Context, in JoinGroupInput) (*entity.GroupMember, error)
}

func NewGroupUseCase(repo groupRepository) GroupUseCaseInterface {
	return &GroupUseCase{
		repo: repo,
	}
}

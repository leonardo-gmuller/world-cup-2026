package schema

import (
	"fmt"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type CreateGroupRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (r *CreateGroupRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}

type JoinGroupRequest struct {
	UserID     int64  `json:"user_id"`
	InviteCode string `json:"invite_code"`
}

func (r *JoinGroupRequest) Validate() error {
	if r.UserID == 0 {
		return fmt.Errorf("user_id is required")
	}

	if r.InviteCode == "" {
		return fmt.Errorf("invite_code is required")
	}

	return nil
}

type GroupResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	InviteCode  string  `json:"invite_code"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
}

func GroupResponseFromEntity(group *entity.Group) GroupResponse {
	return GroupResponse{
		ID:          group.UUID.String(),
		Name:        group.Name,
		Description: group.Description,
		InviteCode:  group.InviteCode,
		IsActive:    group.IsActive,
		CreatedAt:   group.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func GroupListResponseFromEntity(groups []entity.Group) []GroupResponse {
	items := make([]GroupResponse, 0, len(groups))

	for _, group := range groups {
		items = append(items, GroupResponse{
			ID:          group.UUID.String(),
			Name:        group.Name,
			Description: group.Description,
			InviteCode:  group.InviteCode,
			IsActive:    group.IsActive,
			CreatedAt:   group.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return items
}

type GroupMemberResponse struct {
	ID       string `json:"id"`
	GroupID  int64  `json:"group_id"`
	UserID   int64  `json:"user_id"`
	Role     string `json:"role"`
	JoinedAt string `json:"joined_at"`
}

func GroupMemberResponseFromEntity(member *entity.GroupMember) GroupMemberResponse {
	return GroupMemberResponse{
		ID:       member.UUID.String(),
		GroupID:  member.GroupID,
		UserID:   member.UserID,
		Role:     member.Role,
		JoinedAt: member.JoinedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

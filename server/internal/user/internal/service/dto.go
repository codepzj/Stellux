package service

import (
	"server/internal/user/internal/domain"
	"time"

	"github.com/samber/lo"
)

type UserDto struct {
	ID        string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	RoleId    int
}

func DoToDTO(user *domain.User) *UserDto {
	return &UserDto{
		ID:        user.ID.Hex(),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		RoleId:    user.RoleId,
	}
}

func DOsToDTOs(users []*domain.User) []*UserDto {
	return lo.Map(users, func(user *domain.User, _ int) *UserDto {
		return DoToDTO(user)
	})
}

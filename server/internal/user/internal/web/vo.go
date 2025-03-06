package web

import (
	"github.com/codepzj/Stellux/server/internal/user/internal/service"

	"github.com/samber/lo"
)

type UserVO struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	RoleId    *int   `json:"role_id"`
}

type LoginVO struct {
	User  *UserVO `json:"user"`
	Token string  `json:"token"`
}

func DTOToVO(user *service.UserDto) *UserVO {
	return &UserVO{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		RoleId:    user.RoleId,
	}
}

func DTOsToVOs(users []*service.UserDto) []*UserVO {
	return lo.Map(users, func(user *service.UserDto, _ int) *UserVO {
		return DTOToVO(user)
	})
}

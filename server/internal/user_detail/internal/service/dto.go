package service

import "github.com/codepzj/Stellux/server/internal/user_detail/internal/domain"

type UserDetailDTO struct {
	UserId      string
	NickName    string
	Avatar      string
	Description string
	Email       string
}

func DOToDTO(userDetail *domain.UserDetail) *UserDetailDTO {
	return &UserDetailDTO{
		UserId:      userDetail.UserID.Hex(),
		NickName:    userDetail.NickName,
		Avatar:      userDetail.Avatar,
		Description: userDetail.Description,
		Email:       userDetail.Email,
	}
}

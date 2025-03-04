package api

import "server/internal/user_detail/internal/domain"

type UserDetailVO struct {
	UserId      string `json:"user_id" binding:"required"`
	NickName    string `json:"nick_name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Email       string `json:"email"`
}

func DOToVO(userDetail *domain.UserDetail) *UserDetailVO {
	return &UserDetailVO{
		UserId:      userDetail.UserID.Hex(),
		NickName:    userDetail.NickName,
		Avatar:      userDetail.Avatar,
		Description: userDetail.Description,
		Email:       userDetail.Email,
	}
}

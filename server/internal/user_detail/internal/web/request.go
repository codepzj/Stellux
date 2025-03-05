package web

import (
	"server/internal/user_detail/internal/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserDetailRequest struct {
	UserId      string `json:"user_id" binding:"required"`
	NickName    string `json:"nick_name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Email       string `json:"email"`
}

func UserDetailReqToDO(userDetail *UserDetailRequest) *domain.UserDetail {
	userID, err := bson.ObjectIDFromHex(userDetail.UserId)
	if err != nil {
		return nil
	}
	return &domain.UserDetail{
		UserID:      userID,
		NickName:    userDetail.NickName,
		Avatar:      userDetail.Avatar,
		Description: userDetail.Description,
		Email:       userDetail.Email,
	}
}

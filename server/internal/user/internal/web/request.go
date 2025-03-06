package web

import "github.com/codepzj/Stellux/server/internal/user/internal/domain"

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId   *int   `json:"role_id" binding:"required"`
}

func LoginReqToDO(req *LoginReq) *domain.User {
	return &domain.User{
		Username: req.Username,
		Password: req.Password,
	}
}

func CreateUserReqToDO(req *CreateUserReq) *domain.User {
	return &domain.User{
		Username: req.Username,
		Password: req.Password,
		RoleId:   req.RoleId,
	}
}

package user_detail

import (
	"server/internal/user_detail/internal/domain"
	"server/internal/user_detail/internal/service"
	"server/internal/user_detail/internal/web"
)

type (
	Handler    = web.UserDetailHandler
	Service    = service.IUserDetailService
	UserDetail = domain.UserDetail
	Module     struct {
		Hdl *Handler
		Svc Service
	}
)

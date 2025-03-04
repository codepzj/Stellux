package user_detail

import (
	"server/internal/user_detail/internal/api"
	"server/internal/user_detail/internal/service"
)

type (
	Handler = api.UserDetailHandler
	Service = service.IUserDetailService
	Module  struct {
		Hdl *Handler
		Svc Service
	}
)

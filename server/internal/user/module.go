package user

import (
	"server/internal/user/internal/service"
	"server/internal/user/internal/web"
)

type (
	Handler = web.UserHandler
	Service = service.IUserService
	Module  struct {
		Svc Service
		Hdl *Handler
	}
)

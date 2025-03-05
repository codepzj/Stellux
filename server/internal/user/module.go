package user

import (
	"server/internal/user/internal/repo"
	"server/internal/user/internal/service"
	"server/internal/user/internal/web"
)

type (
	Handler    = web.UserHandler
	Service    = service.IUserService
	Repository = repo.IUserRepo
	Module     struct {
		Hdl *Handler
		Svc  Service
		Repo Repository
	}
)

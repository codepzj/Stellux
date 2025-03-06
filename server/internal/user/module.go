package user

import (
	"github.com/codepzj/Stellux/server/internal/user/internal/repo"
	"github.com/codepzj/Stellux/server/internal/user/internal/service"
	"github.com/codepzj/Stellux/server/internal/user/internal/web"
)

type (
	Handler    = web.UserHandler
	Service    = service.IUserService
	Repository = repo.IUserRepo
	Module     struct {
		Hdl  *Handler
		Svc  Service
		Repo Repository
	}
)

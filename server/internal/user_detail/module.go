package user_detail

import (
	"github.com/codepzj/Stellux/server/internal/user_detail/internal/domain"
	"github.com/codepzj/Stellux/server/internal/user_detail/internal/repo"
	"github.com/codepzj/Stellux/server/internal/user_detail/internal/service"
	"github.com/codepzj/Stellux/server/internal/user_detail/internal/web"
)

type (
	Handler    = web.UserDetailHandler
	Service    = service.IUserDetailService
	Repository = repo.IUserDetailRepo
	UserDetail = domain.UserDetail
	Module     struct {
		Hdl  *Handler
		Svc  Service
		Repo Repository
	}
)

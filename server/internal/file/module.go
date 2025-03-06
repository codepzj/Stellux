package file

import (
	"github.com/codepzj/Stellux/server/internal/file/internal/repo"
	"github.com/codepzj/Stellux/server/internal/file/internal/service"
	"github.com/codepzj/Stellux/server/internal/file/internal/web"
)

type (
	Handler    = web.FileHandler
	Service    = service.IFileService
	Repository = repo.IFileRepo
	Module     struct {
		Hdl  *Handler
		Svc  Service
		Repo Repository
	}
)

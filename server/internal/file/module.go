package file

import (
	"server/internal/file/internal/repo"
	"server/internal/file/internal/service"
	"server/internal/file/internal/web"
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

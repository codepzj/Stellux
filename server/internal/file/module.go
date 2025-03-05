package file

import (
	"server/internal/file/internal/service"
	"server/internal/file/internal/web"
)

type (
	Handler = web.FileHandler
	Service = service.IFileService
	Module  struct {
		Hdl *Handler
		Svc Service
	}
)

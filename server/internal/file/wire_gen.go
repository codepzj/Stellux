// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package file

import (
	"github.com/google/wire"
	"github.com/codepzj/Stellux/server/internal/file/internal/repo"
	"github.com/codepzj/Stellux/server/internal/file/internal/repo/dao"
	"github.com/codepzj/Stellux/server/internal/file/internal/service"
	"github.com/codepzj/Stellux/server/internal/file/internal/web"
)

// Injectors from wire.go:

func InitFileModule() *Module {
	fileDao := dao.NewFileDao()
	fileRepo := repo.NewFileRepo(fileDao)
	fileService := service.NewFileService(fileRepo)
	fileHandler := web.NewFileHandler(fileService)
	module := &Module{
		Hdl:  fileHandler,
		Svc:  fileService,
		Repo: fileRepo,
	}
	return module
}

// wire.go:

var fileProvider = wire.NewSet(web.NewFileHandler, service.NewFileService, repo.NewFileRepo, dao.NewFileDao, wire.Bind(new(service.IFileService), new(*service.FileService)), wire.Bind(new(repo.IFileRepo), new(*repo.FileRepo)), wire.Bind(new(dao.IFileDao), new(*dao.FileDao)))

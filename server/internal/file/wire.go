//go:build wireinject
// +build wireinject

package file

import (
	"github.com/codepzj/Stellux/server/internal/file/internal/repo"
	"github.com/codepzj/Stellux/server/internal/file/internal/repo/dao"
	"github.com/codepzj/Stellux/server/internal/file/internal/service"
	"github.com/codepzj/Stellux/server/internal/file/internal/web"

	"github.com/google/wire"
)

var fileProvider = wire.NewSet(
	web.NewFileHandler,
	service.NewFileService,
	repo.NewFileRepo,
	dao.NewFileDao,
	wire.Bind(new(service.IFileService), new(*service.FileService)),
	wire.Bind(new(repo.IFileRepo), new(*repo.FileRepo)),
	wire.Bind(new(dao.IFileDao), new(*dao.FileDao)),
)

func InitFileModule() *Module {
	wire.Build(
		fileProvider,
		wire.Struct(new(Module), "*"),
	)
	return nil
}

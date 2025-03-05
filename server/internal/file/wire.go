//go:build wireinject
// +build wireinject

package file

import (
	"server/internal/file/internal/repo"
	"server/internal/file/internal/repo/dao"
	"server/internal/file/internal/service"
	"server/internal/file/internal/web"

	"go.mongodb.org/mongo-driver/v2/mongo"

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

func InitFileModule(database *mongo.Database) *Module {
	wire.Build(
		fileProvider,
		wire.Struct(new(Module), "Hdl", "Svc"),
	)
	return nil
}

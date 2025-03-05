//go:build wireinject
// +build wireinject

package user_detail

import (
	"server/internal/user_detail/internal/repo"
	"server/internal/user_detail/internal/repo/dao"
	"server/internal/user_detail/internal/service"
	"server/internal/user_detail/internal/web"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userDetailProvider = wire.NewSet(
	web.NewUserDetailHandler, service.NewUserDetailService, repo.NewUserDetailRepo, dao.NewUserDetailDao,
	wire.Bind(new(service.IUserDetailService), new(*service.UserDetailService)),
	wire.Bind(new(repo.IUserDetailRepo), new(*repo.UserDetailRepo)),
	wire.Bind(new(dao.IUserDetailDao), new(*dao.UserDetailDao)),
)

func InitUserDetailModule(db *mongo.Database) *Module {
	wire.Build(
		userDetailProvider,
		wire.Struct(new(Module), "Hdl", "Svc"),
	)
	return nil
}

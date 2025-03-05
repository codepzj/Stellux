//go:build wireinject
// +build wireinject

package user

import (
	"server/internal/user/internal/repo"
	"server/internal/user/internal/repo/dao"
	"server/internal/user/internal/service"
	"server/internal/user/internal/web"
	"server/internal/user_detail"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userProvider = wire.NewSet(web.NewUserHandler, service.NewUserService, repo.NewUserRepo, dao.NewUserDao,
	wire.Bind(new(service.IUserService), new(*service.UserService)),
	wire.Bind(new(repo.IUserRepo), new(*repo.UserRepo)),
	wire.Bind(new(dao.IUserDao), new(*dao.UserDao)))

func InitUserModule(database *mongo.Database, userDetailModule *user_detail.Module) *Module {
	panic(wire.Build(
		userProvider,
		wire.FieldsOf(new(*user_detail.Module), "Svc"),
		wire.Struct(new(Module), "Hdl", "Svc"),
	))
}

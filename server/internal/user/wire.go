//go:build wireinject
// +build wireinject

package user

import (
	"github.com/codepzj/Stellux/server/internal/user/internal/repo"
	"github.com/codepzj/Stellux/server/internal/user/internal/repo/dao"
	"github.com/codepzj/Stellux/server/internal/user/internal/service"
	"github.com/codepzj/Stellux/server/internal/user/internal/web"
	"github.com/codepzj/Stellux/server/internal/user_detail"

	"github.com/google/wire"
)

var userProvider = wire.NewSet(web.NewUserHandler, service.NewUserService, repo.NewUserRepo, dao.NewUserDao,
	wire.Bind(new(service.IUserService), new(*service.UserService)),
	wire.Bind(new(repo.IUserRepo), new(*repo.UserRepo)),
	wire.Bind(new(dao.IUserDao), new(*dao.UserDao)))

func InitUserModule(userDetailModule *user_detail.Module) *Module {
	panic(wire.Build(
		userProvider,
		wire.FieldsOf(new(*user_detail.Module), "Repo"),
		wire.Struct(new(Module), "*"),
	))
}

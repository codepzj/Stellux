// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user_detail

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/internal/user_detail/internal/api"
	"server/internal/user_detail/internal/repo"
	"server/internal/user_detail/internal/repo/dao"
	"server/internal/user_detail/internal/service"
)

// Injectors from wire.go:

func InitUserDetailModule(db *mongo.Database) *Module {
	userDetailDao := dao.NewUserDetailDao(db)
	userDetailRepo := repo.NewUserDetailRepo(userDetailDao)
	userDetailService := service.NewUserDetailService(userDetailRepo)
	userDetailHandler := api.NewUserDetailHandler(userDetailService)
	module := &Module{
		Hdl: userDetailHandler,
		Svc: userDetailService,
	}
	return module
}

// wire.go:

var userDetailProvider = wire.NewSet(api.NewUserDetailHandler, service.NewUserDetailService, repo.NewUserDetailRepo, dao.NewUserDetailDao, wire.Bind(new(service.IUserDetailService), new(*service.UserDetailService)), wire.Bind(new(repo.IUserDetailRepo), new(*repo.UserDetailRepo)), wire.Bind(new(dao.IUserDetailDao), new(*dao.UserDetailDao)))

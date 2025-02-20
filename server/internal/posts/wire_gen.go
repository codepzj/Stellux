// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package posts

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/google/wire"
	"server/internal/posts/internal/api"
	"server/internal/posts/internal/repo"
	"server/internal/posts/internal/repo/dao"
	"server/internal/posts/internal/service"
)

// Injectors from wire.go:

func InitPostsModule(database *mongox.Database) *Module {
	postsDao := dao.NewPostsDao(database)
	postsRepo := repo.NewPostsRepo(postsDao)
	postsService := service.NewPostsService(postsRepo)
	postsHandler := api.NewPostHandler(postsService)
	module := &Module{
		Hdl: postsHandler,
		Svc: postsService,
	}
	return module
}

// wire.go:

var postsProvider = wire.NewSet(api.NewPostHandler, service.NewPostsService, repo.NewPostsRepo, dao.NewPostsDao, wire.Bind(new(service.IPostsService), new(*service.PostsService)), wire.Bind(new(repo.IPostsRepo), new(*repo.PostsRepo)), wire.Bind(new(dao.IPostsDao), new(*dao.PostsDao)), wire.Struct(new(Module), "Hdl", "Svc"))

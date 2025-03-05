// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package posts

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/internal/posts/internal/repo"
	"server/internal/posts/internal/repo/dao"
	"server/internal/posts/internal/service"
	"server/internal/posts/internal/web"
)

// Injectors from wire.go:

func InitPostsModule(database *mongo.Database) *Module {
	postsDao := dao.NewPostsDao(database)
	postsRepo := repo.NewPostsRepo(postsDao)
	postsService := service.NewPostsService(postsRepo)
	postsHandler := web.NewPostHandler(postsService)
	module := &Module{
		Hdl: postsHandler,
		Svc: postsService,
	}
	return module
}

// wire.go:

var postsProvider = wire.NewSet(web.NewPostHandler, service.NewPostsService, repo.NewPostsRepo, dao.NewPostsDao, wire.Bind(new(service.IPostsService), new(*service.PostsService)), wire.Bind(new(repo.IPostsRepo), new(*repo.PostsRepo)), wire.Bind(new(dao.IPostsDao), new(*dao.PostsDao)))

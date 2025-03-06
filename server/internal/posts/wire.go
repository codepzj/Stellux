//go:build wireinject
// +build wireinject

package posts

import (
	"github.com/codepzj/Stellux/server/internal/posts/internal/repo"
	"github.com/codepzj/Stellux/server/internal/posts/internal/repo/dao"
	"github.com/codepzj/Stellux/server/internal/posts/internal/service"
	"github.com/codepzj/Stellux/server/internal/posts/internal/web"

	"github.com/google/wire"
)

var postsProvider = wire.NewSet(
	web.NewPostHandler,
	service.NewPostsService,
	repo.NewPostsRepo,
	dao.NewPostsDao,
	wire.Bind(new(service.IPostsService), new(*service.PostsService)),
	wire.Bind(new(repo.IPostsRepo), new(*repo.PostsRepo)),
	wire.Bind(new(dao.IPostsDao), new(*dao.PostsDao)),
)

func InitPostsModule() *Module {
	wire.Build(
		postsProvider,
		wire.Struct(new(Module), "*"),
	)
	return nil
}

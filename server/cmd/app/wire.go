//go:build wireinject
// +build wireinject

package app

import (
	"github.com/codepzj/Stellux/server/internal/ioc"

	"github.com/codepzj/Stellux/server/internal/file"
	"github.com/codepzj/Stellux/server/internal/posts"
	"github.com/codepzj/Stellux/server/internal/user"
	"github.com/codepzj/Stellux/server/internal/user_detail"

	"github.com/google/wire"
)

func InitApp() *HttpServer {
	wire.Build(
		ioc.InitMiddleWare,
		ioc.NewGin,

		user.InitUserModule,
		wire.FieldsOf(new(*user.Module), "Hdl"),

		user_detail.InitUserDetailModule,
		wire.FieldsOf(new(*user_detail.Module), "Hdl"),

		posts.InitPostsModule,
		wire.FieldsOf(new(*posts.Module), "Hdl"),

		file.InitFileModule,
		wire.FieldsOf(new(*file.Module), "Hdl"),

		NewHttpServer,
	)
	return nil
}

//go:build wireinject
// +build wireinject

package main

import (
	"server/internal/ioc"

	"server/internal/file"
	"server/internal/posts"
	"server/internal/user"
	"server/internal/user_detail"

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

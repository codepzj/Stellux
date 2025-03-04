//go:build wireinject
// +build wireinject

package main

import (
	"server/internal/file"
	"server/internal/ioc"
	"server/internal/posts"
	"server/internal/user"
	"server/internal/user_detail"

	"github.com/google/wire"
)

func InitApp() *HttpServer {
	wire.Build(
		ioc.NewMongoDB,
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

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"server/internal/file"
	"server/internal/ioc"
	"server/internal/posts"
	"server/internal/user"
	"server/internal/user_detail"
)

// Injectors from wire.go:

func InitApp() *HttpServer {
	database := ioc.NewMongoDB()
	module := user.InitUserModule(database)
	userHandler := module.Hdl
	user_detailModule := user_detail.InitUserDetailModule(database)
	userDetailHandler := user_detailModule.Hdl
	postsModule := posts.InitPostsModule(database)
	postsHandler := postsModule.Hdl
	fileModule := file.InitFileModule(database)
	fileHandler := fileModule.Hdl
	v := ioc.InitMiddleWare()
	engine := ioc.NewGin(userHandler, userDetailHandler, postsHandler, fileHandler, v)
	httpServer := NewHttpServer(engine)
	return httpServer
}

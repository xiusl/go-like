// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go-like/app/interface/service/internal/biz"
	"go-like/app/interface/service/internal/conf"
	"go-like/app/interface/service/internal/data"
	"go-like/app/interface/service/internal/server"
	"go-like/app/interface/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserClient(discovery)
	dataData := data.NewData(confData, userClient, logger)
	userRepo := data.NewUserRepo(dataData)
	authUseCase := biz.NewAuthUseCase(userRepo, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	interfaceService := service.NewInterfaceService(logger, authUseCase, userUseCase)
	httpServer := server.NewHTTPServer(confServer, interfaceService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
	}, nil
}

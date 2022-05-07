// +build wireinject

package main

import (
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    "go-like/app/interface/service/internal/biz"
    "go-like/app/interface/service/internal/conf"
    "go-like/app/interface/service/internal/data"
    "go-like/app/interface/service/internal/server"
    "go-like/app/interface/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
    panic(wire.Build(
        server.ProviderSet,
        service.ProviderSet,
        biz.ProviderSet,
        data.ProviderSet,
        newApp,
    ))
}

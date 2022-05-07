package service

import (
    "github.com/gin-gonic/gin"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewInterfaceService)

func NewInterfaceService(logger log.Logger) *InterfaceService {
    return &InterfaceService{
        log: log.NewHelper(log.With(logger, "module", "interface/service")),
    }
}

type InterfaceService struct {
    *gin.Engine
    log *log.Helper
}

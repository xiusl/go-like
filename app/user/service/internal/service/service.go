package service

import (
    "github.com/gin-gonic/gin"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    v1 "go-like/api/user/service/v1"
)

var ProviderSet = wire.NewSet(NewUserService)

func NewUserService(logger log.Logger) *UserService {
	return &UserService{
		log:          log.NewHelper(log.With(logger, "module", "user/service")),
	}
}

type UserService struct {
    v1.UnimplementedUserServer
    *gin.Engine
    log *log.Helper
}
package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "go-like/api/user/service/v1"
	"go-like/app/user/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewUserService)

func NewUserService(authUc *biz.AuthUseCase, userUc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		log:    log.NewHelper(log.With(logger, "module", "user/service")),
		userUc: userUc,
		authUc: authUc,
	}
}

type UserService struct {
	v1.UnimplementedUserServer
	log *log.Helper

	authUc *biz.AuthUseCase
	userUc *biz.UserUseCase
}

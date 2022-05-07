package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-like/app/interface/service/internal/biz"
	"go-like/pkg/xgin"
)

var ProviderSet = wire.NewSet(NewInterfaceService)

type InterfaceService struct {
	*gin.Engine
	log    *log.Helper
	authUc *biz.AuthUseCase
	userUc *biz.UserUseCase
}

func NewInterfaceService(
	logger log.Logger,
	authUc *biz.AuthUseCase,
	userUc *biz.UserUseCase,
) *InterfaceService {

	engine := gin.New()
	engine.Use(xgin.Logger(logger), gin.CustomRecovery(xgin.XRecoveryHandler()))

	srv := &InterfaceService{
		Engine: engine,
		log:    log.NewHelper(log.With(logger, "module", "interface/service")),
		authUc: authUc,
		userUc: userUc,
	}
	srv.mapper()
	return srv
}

func (srv *InterfaceService) mapper() {
	r := srv.Group("")
	xgin.GET(r, "/ping", srv.pong)
	srv.mapAuth()
	srv.mapUser()
	srv.mapVerifyCode()
}

func (srv *InterfaceService) pong(ctx *gin.Context) (interface{}, string, int) {
	return "pong", "", 200
}

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
	artUc  *biz.ArticleUseCase
}

func NewInterfaceService(
	logger log.Logger,
	authUc *biz.AuthUseCase,
	userUc *biz.UserUseCase,
	artUc *biz.ArticleUseCase,
) *InterfaceService {

	engine := gin.New()
	engine.Use(xgin.Logger(logger))
	engine.Use(gin.CustomRecovery(xgin.XRecoveryHandler()))
	engine.Use(xgin.UserMiddleware())

	srv := &InterfaceService{
		Engine: engine,
		log:    log.NewHelper(log.With(logger, "module", "interface/service")),
		authUc: authUc,
		userUc: userUc,
		artUc:  artUc,
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

	srv.mapArticle()
}

func (srv *InterfaceService) pong(ctx *gin.Context) (interface{}, string, int) {
	return "pong", "", 200
}

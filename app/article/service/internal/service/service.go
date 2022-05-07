package service

import (
    "github.com/gin-gonic/gin"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    v1 "go-like/api/article/service/v1"
)

var ProviderSet = wire.NewSet(NewArticleService)

func NewArticleService(logger log.Logger) *ArticleService {
	return &ArticleService{
		log:          log.NewHelper(log.With(logger, "module", "article/service")),
	}
}

type ArticleService struct {
    v1.UnimplementedArticleServer
    *gin.Engine
    log *log.Helper
}
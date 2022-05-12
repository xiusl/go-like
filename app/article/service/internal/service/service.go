package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "go-like/api/article/service/v1"
	"go-like/app/article/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewArticleService)

func NewArticleService(logger log.Logger, artUc *biz.ArticleUseCase) *ArticleService {
	return &ArticleService{
		artUc: artUc,
		log:   log.NewHelper(log.With(logger, "module", "article/service")),
	}
}

type ArticleService struct {
	v1.UnimplementedArticleServer

	log   *log.Helper
	artUc *biz.ArticleUseCase
}

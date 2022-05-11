package data

import (
	"context"
	artv1 "go-like/api/article/service/v1"
	"go-like/app/interface/service/internal/biz"
)

type articleRepo struct {
	data *Data
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleRepo{
		data: data,
	}
}

func parseGrpcArticle(art *artv1.ArticleInfo) *biz.Article {
	if art == nil {
		return nil
	}
	return &biz.Article{
		Id:     art.Id,
		Title:  art.Title,
		Url:    art.Url,
		UserId: art.UserId,
	}
}

func (r *articleRepo) CreateArticle(ctx context.Context, uid int64, title, url string) (*biz.Article, error) {
	rep, err := r.data.ac.PublishArticle(ctx, &artv1.PublishArticleReq{
		UserId: uid,
		Title:  title,
		Url:    url,
	})
	if err != nil {
		return nil, err
	}
	return parseGrpcArticle(rep.Article), nil
}

func (r *articleRepo) GetArticle(ctx context.Context, aid int64) (*biz.Article, error) {
	rep, err := r.data.ac.GetArticle(ctx, &artv1.ArticleIdReq{
		Id: aid,
	})
	if err != nil {
		return nil, err
	}
	return parseGrpcArticle(rep.Article), nil
}

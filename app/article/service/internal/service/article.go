package service

import (
	"context"
	v1 "go-like/api/article/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *ArticleService) PublishArticle(ctx context.Context, req *v1.PublishArticleReq) (*v1.ArticleInfoReply, error) {
	art, err := srv.artUc.CreateArticle(ctx, req.Title, req.Url, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.ArticleInfoReply{
		Article: &v1.ArticleInfo{
			Id:     art.Id,
			Title:  art.Title,
			Url:    art.Url,
			UserId: art.UserId,
		},
	}, nil
}
func (srv *ArticleService) GetArticle(ctx context.Context, req *v1.ArticleIdReq) (*v1.ArticleInfoReply, error) {
	art, err := srv.artUc.GetArticle(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ArticleInfoReply{
		Article: &v1.ArticleInfo{
			Id:     art.Id,
			Title:  art.Title,
			Url:    art.Url,
			UserId: art.UserId,
		},
	}, nil
}
func (srv *ArticleService) ListArticles(ctx context.Context, req *v1.PageReq) (*v1.ArticleInfosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}

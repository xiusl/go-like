package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	Insert(ctx context.Context, art *Article) (int64, error)
	Select(ctx context.Context, id int64) (*Article, error)
}

type Article struct {
	Id        int64
	Title     string
	Url       string
	Content   string
	Images    []string
	UserId    int64
	CreatedAt time.Time
}

type ArticleUseCase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUseCase(repo ArticleRepo, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ArticleUseCase) CreateArticle(ctx context.Context, title, url string, uid int64) (*Article, error) {
	id, err := uc.repo.Insert(ctx, &Article{
		Title:  title,
		Url:    url,
		UserId: uid,
	})

	if err != nil {
		return nil, err
	}

	art, err := uc.repo.Select(ctx, id)
	if err != nil {
		return nil, err
	}

	// MQ get content

	return art, nil
}

func (uc *ArticleUseCase) GetArticle(ctx context.Context, id int64) (*Article, error) {
	art, err := uc.repo.Select(ctx, id)
	if err != nil {
		return nil, err
	}
	return art, nil
}

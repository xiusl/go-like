package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	CreateArticle(ctx context.Context, uid int64, title, url string) (*Article, error)
	GetArticle(ctx context.Context, aid int64) (*Article, error)
}

type Article struct {
	Id     int64
	Title  string
	Url    string
	UserId int64
}

type ArticleUseCase struct {
	repo     ArticleRepo
	userRepo UserRepo
	log      *log.Helper
}

func NewArticleUseCase(repo ArticleRepo, userRepo UserRepo, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		repo:     repo,
		userRepo: userRepo,
		log:      log.NewHelper(logger),
	}
}

func (uc *ArticleUseCase) CreateArticle(ctx context.Context, uid int64, title, url string) (*Article, *User, error) {
	art, err := uc.repo.CreateArticle(ctx, uid, title, url)
	if err != nil {
		return nil, nil, err
	}
	user, err := uc.userRepo.GetUser(ctx, uid)
	if err != nil {
		return nil, nil, err
	}
	return art, user, nil
}

func (uc *ArticleUseCase) GetArticle(ctx context.Context, uid int64, aid int64) (*Article, *User, error) {
	art, err := uc.repo.GetArticle(ctx, aid)
	if err != nil {
		return nil, nil, err
	}
	user, err := uc.userRepo.GetUser(ctx, art.UserId)
	if err != nil {
		return nil, nil, err
	}
	if uid > 0 {
		// feat: 获取是否点赞，是否关注等信息
		uc.log.Info("current user", uid)
	}

	return art, user, nil
}

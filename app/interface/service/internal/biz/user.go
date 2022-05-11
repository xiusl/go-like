package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	Auth(ctx context.Context, mobile, verifyCode string) (*User, string, error)
	VerifyToken(ctx context.Context, token string) (int64, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	SendVerifyCode(ctx context.Context, key string, bizType int64) error
}

type User struct {
	Id     int64
	Name   string
	Mobile string
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) SendVerifyCode(ctx context.Context, key string, bizType int64) error {
	return uc.repo.SendVerifyCode(ctx, key, bizType)
}

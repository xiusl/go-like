package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewAuthUseCase(repo UserRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *AuthUseCase) Auth(ctx context.Context, mobile, verifyCode string) (*User, string, error) {
	return uc.repo.Auth(ctx, mobile, verifyCode)
}

func (uc *AuthUseCase) VerifyToken(ctx context.Context, token string) (int64, error) {
	return uc.repo.VerifyToken(ctx, token)
}

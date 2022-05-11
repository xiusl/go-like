package biz

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"go-like/pkg/security"
	"time"
)

type AuthUseCase struct {
	uRepo  UserRepo
	vcRepo VerifyCodeRepo
	log    *log.Helper
}

func NewAuthUseCase(repo UserRepo, vcRepo VerifyCodeRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		uRepo:  repo,
		vcRepo: vcRepo,
		log:    log.NewHelper(logger),
	}
}

func (uc *AuthUseCase) Auth(ctx context.Context, mobile, code string) (*User, string, error) {
	c, err := uc.vcRepo.SelectByKey(ctx, mobile, 0)
	if err != nil {
		return nil, "", err
	}

	if c.Code != code || c.ExpiredAt.Before(time.Now()) {
		return nil, "", errors.New("无效的验证码")
	}

	u, err := uc.uRepo.SelectByMobile(ctx, mobile)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, "", err
		}

		u = &User{
			Mobile: mobile,
		}

		u.Id, err = uc.uRepo.Insert(ctx, u)
		if err != nil {
			return nil, "", err
		}
	}

	token := security.GenerateToken(u.Id, "")
	return u, token, nil
}

func (uc *AuthUseCase) VerifyToken(ctx context.Context, token string) (int64, error) {
	uid, sign, msg, err := security.ParseToken(token)
	if err != nil {
		return 0, err
	}
	if err := security.VerifyTokenSign(sign, msg, ""); err != nil {
		return 0, err
	}
	return uid, nil
}

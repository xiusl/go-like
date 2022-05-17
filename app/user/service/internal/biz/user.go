package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type UserRepo interface {
	Select(ctx context.Context, id int64) (*User, error)
	SelectByMobile(ctx context.Context, mobile string) (*User, error)
	Insert(ctx context.Context, u *User) (int64, error)
	UpdateFollowerCount(ctx context.Context, uid int64, count int64) error
	ListByIds(ctx context.Context, ids []int64) ([]*User, error)
	MapByIds(ctx context.Context, ids []int64) (map[int64]*User, error)
}

type User struct {
	Id             int64
	Name           string
	Mobile         string
	Avatar         string
	FollowerCount  int64
	FollowingCount int64
	IsFollowed     bool
	IsFollowing    bool
}

type UserUseCase struct {
	repo   UserRepo
	vcRepo VerifyCodeRepo
	fRepo  FollowerRepo
	log    *log.Helper
}

func NewUserUseCase(repo UserRepo, vcRepo VerifyCodeRepo, fRepo FollowerRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		vcRepo: vcRepo,
		fRepo:  fRepo,
		log:    log.NewHelper(logger),
	}
}

func (uc *UserUseCase) SendVerifyCode(ctx context.Context, key string, bizType int64) error {
	// todo: 控制频次
	return uc.vcRepo.Insert(ctx, &VerifyCode{
		Key:       key,
		BizType:   bizType,
		Code:      "000000", // todo: 随机数生成
		ExpiredAt: time.Now().Add(time.Minute * 10),
	})
}

func (uc *UserUseCase) GetUser(ctx context.Context, uid int64) (*User, error) {
	u, err := uc.repo.Select(ctx, uid)
	if err != nil {
		return nil, err
	}
	return u, nil
}

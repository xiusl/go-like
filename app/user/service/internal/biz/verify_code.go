package biz

import (
	"context"
	"time"
)

type VerifyCodeRepo interface {
	Select(ctx context.Context, id int64) (*VerifyCode, error)
	SelectByKey(ctx context.Context, key string, bizType int64) (*VerifyCode, error)
	Insert(ctx context.Context, c *VerifyCode) error
}

type VerifyCode struct {
	Id        int64
	Key       string
	Code      string
	BizType   int64
	ExpiredAt time.Time
}

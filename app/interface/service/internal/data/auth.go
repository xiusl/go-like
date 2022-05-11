package data

import (
	"context"
	v1 "go-like/api/user/service/v1"
)

func (u *userRepo) VerifyToken(ctx context.Context, token string) (int64, error) {
	rep, err := u.data.uc.VerifyToken(ctx, &v1.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		return 0, err
	}
	return rep.UserId, nil
}

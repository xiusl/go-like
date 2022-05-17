package data

import (
	"context"
	v1 "go-like/api/user/service/v1"
	"go-like/app/interface/service/internal/biz"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func parseGrpcUser(user *v1.UserInfo) *biz.User {
	if user == nil {
		return nil
	}
	return &biz.User{
		Id:          user.Id,
		Name:        user.Name,
		Mobile:      user.Mobile,
		Avatar:      user.Avatar,
		IsFollowing: user.IsFollowing,
		IsFollowed:  user.IsFollowed,
	}
}

func (u *userRepo) Auth(ctx context.Context, mobile, verifyCode string) (*biz.User, string, error) {
	rep, err := u.data.uc.Auth(ctx, &v1.AuthReq{
		Mobile: mobile,
		Code:   verifyCode,
	})
	if err != nil {
		return nil, "", err
	}
	return parseGrpcUser(rep.User), rep.Token, nil
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	rep, err := u.data.uc.GetUser(ctx, &v1.GetUserReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return parseGrpcUser(rep.User), nil
}

func (u *userRepo) SendVerifyCode(ctx context.Context, key string, bizType int64) error {
	_, err := u.data.uc.SendVerifyCode(ctx, &v1.SendVerifyCodeReq{
		Key:     key,
		BizType: bizType,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) FollowingUser(ctx context.Context, authUser, id int64) error {
	_, err := u.data.uc.FollowingUser(ctx, &v1.FollowingUserReq{
		User:      authUser,
		Following: id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetFollowings(ctx context.Context, uid int64) ([]*biz.User, error) {
	rep, err := u.data.uc.GetFollowings(ctx, &v1.UserIdPageReq{
		Uid: uid,
		Page: &v1.PageReq{
			Page:  1,
			Count: 10,
		},
	})
	if err != nil {
		return nil, err
	}
	us := make([]*biz.User, 0)
	for _, rpcUser := range rep.Users {
		us = append(us, parseGrpcUser(rpcUser))
	}
	return us, nil
}

func (u *userRepo) GetFollowers(ctx context.Context, uid int64) ([]*biz.User, error) {
	rep, err := u.data.uc.GetFollowers(ctx, &v1.UserIdPageReq{
		Uid: uid,
		Page: &v1.PageReq{
			Page:  1,
			Count: 10,
		},
	})
	if err != nil {
		return nil, err
	}
	us := make([]*biz.User, 0)
	for _, rpcUser := range rep.Users {
		us = append(us, parseGrpcUser(rpcUser))
	}
	return us, nil
}

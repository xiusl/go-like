package service

import (
	"context"
	v1 "go-like/api/user/service/v1"
	"go-like/app/user/service/internal/biz"
)

func (srv *UserService) Auth(ctx context.Context, req *v1.AuthReq) (*v1.AuthReply, error) {
	u, token, err := srv.authUc.Auth(ctx, req.Mobile, req.Code)
	if err != nil {
		return nil, err
	}
	return &v1.AuthReply{
		User: &v1.UserInfo{
			Id:     u.Id,
			Name:   u.Name,
			Mobile: u.Mobile,
		},
		Token: token,
	}, nil
}
func (srv *UserService) SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeReq) (*v1.BoolReply, error) {
	err := srv.userUc.SendVerifyCode(ctx, req.Key, req.BizType)
	if err != nil {
		return nil, err
	}
	return &v1.BoolReply{
		Success: true,
	}, nil
}
func (srv *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	u, err := srv.userUc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		User: &v1.UserInfo{
			Id:     u.Id,
			Name:   u.Name,
			Mobile: u.Mobile,
		},
	}, nil
}
func (srv *UserService) VerifyToken(ctx context.Context, req *v1.VerifyTokenReq) (*v1.VerifyTokenReply, error) {
	id, err := srv.authUc.VerifyToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &v1.VerifyTokenReply{
		UserId: id,
	}, nil
}

func (srv *UserService) FollowingUser(ctx context.Context, req *v1.FollowingUserReq) (*v1.BoolReply, error) {
	err := srv.userUc.FollowUser(ctx, req.Following, req.User)
	if err != nil {
		return nil, err
	}
	return &v1.BoolReply{Success: true}, nil
}

func (srv *UserService) GetFollowings(ctx context.Context, req *v1.UserIdPageReq) (*v1.UsersReq, error) {
	users, err := srv.userUc.GetFollowing(ctx, req.Uid, req.Page.Page, req.Page.Count)
	if err != nil {
		return nil, err
	}
	return &v1.UsersReq{
		Users: packUsers(users),
	}, nil
}
func (srv *UserService) GetFollowers(ctx context.Context, req *v1.UserIdPageReq) (*v1.UsersReq, error) {
	users, err := srv.userUc.GetFollower(ctx, req.Uid, req.Page.Page, req.Page.Count)
	if err != nil {
		return nil, err
	}
	return &v1.UsersReq{
		Users: packUsers(users),
	}, nil
}

func packUser(u *biz.User) *v1.UserInfo {
	if u == nil {
		return nil
	}
	return &v1.UserInfo{
		Id:             u.Id,
		Name:           u.Name,
		Mobile:         u.Mobile,
		Avatar:         u.Avatar,
		FollowerCount:  u.FollowerCount,
		FollowingCount: u.FollowingCount,
		IsFollowed:     u.IsFollowed,
		IsFollowing:    u.IsFollowing,
	}
}
func packUsers(us []*biz.User) []*v1.UserInfo {
	res := make([]*v1.UserInfo, len(us))
	for i, u := range us {
		res[i] = packUser(u)
	}
	return res
}

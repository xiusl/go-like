package service

import (
	"context"
	v1 "go-like/api/user/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

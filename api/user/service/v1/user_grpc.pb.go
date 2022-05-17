// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error)
	VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenReply, error)
	SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*BoolReply, error)
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserReply, error)
	FollowingUser(ctx context.Context, in *FollowingUserReq, opts ...grpc.CallOption) (*BoolReply, error)
	GetFollowings(ctx context.Context, in *UserIdPageReq, opts ...grpc.CallOption) (*UsersReq, error)
	GetFollowers(ctx context.Context, in *UserIdPageReq, opts ...grpc.CallOption) (*UsersReq, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/user.service.v1.User/auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenReply, error) {
	out := new(VerifyTokenReply)
	err := c.cc.Invoke(ctx, "/user.service.v1.User/verifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/user.service.v1.User/sendVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/user.service.v1.User/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FollowingUser(ctx context.Context, in *FollowingUserReq, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/user.service.v1.User/followingUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetFollowings(ctx context.Context, in *UserIdPageReq, opts ...grpc.CallOption) (*UsersReq, error) {
	out := new(UsersReq)
	err := c.cc.Invoke(ctx, "/user.service.v1.User/getFollowings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetFollowers(ctx context.Context, in *UserIdPageReq, opts ...grpc.CallOption) (*UsersReq, error) {
	out := new(UsersReq)
	err := c.cc.Invoke(ctx, "/user.service.v1.User/getFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Auth(context.Context, *AuthReq) (*AuthReply, error)
	VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenReply, error)
	SendVerifyCode(context.Context, *SendVerifyCodeReq) (*BoolReply, error)
	GetUser(context.Context, *GetUserReq) (*GetUserReply, error)
	FollowingUser(context.Context, *FollowingUserReq) (*BoolReply, error)
	GetFollowings(context.Context, *UserIdPageReq) (*UsersReq, error)
	GetFollowers(context.Context, *UserIdPageReq) (*UsersReq, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Auth(context.Context, *AuthReq) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedUserServer) VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedUserServer) SendVerifyCode(context.Context, *SendVerifyCodeReq) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerifyCode not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *GetUserReq) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) FollowingUser(context.Context, *FollowingUserReq) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowingUser not implemented")
}
func (UnimplementedUserServer) GetFollowings(context.Context, *UserIdPageReq) (*UsersReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowings not implemented")
}
func (UnimplementedUserServer) GetFollowers(context.Context, *UserIdPageReq) (*UsersReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.User/auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.User/verifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerifyToken(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerifyCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.User/sendVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendVerifyCode(ctx, req.(*SendVerifyCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.User/getUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FollowingUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowingUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FollowingUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.User/followingUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FollowingUser(ctx, req.(*FollowingUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetFollowings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetFollowings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.User/getFollowings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetFollowings(ctx, req.(*UserIdPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.User/getFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetFollowers(ctx, req.(*UserIdPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "auth",
			Handler:    _User_Auth_Handler,
		},
		{
			MethodName: "verifyToken",
			Handler:    _User_VerifyToken_Handler,
		},
		{
			MethodName: "sendVerifyCode",
			Handler:    _User_SendVerifyCode_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "followingUser",
			Handler:    _User_FollowingUser_Handler,
		},
		{
			MethodName: "getFollowings",
			Handler:    _User_GetFollowings_Handler,
		},
		{
			MethodName: "getFollowers",
			Handler:    _User_GetFollowers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/user.proto",
}

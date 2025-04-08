// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: user/service/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	User_RegisterWithCredentials_FullMethodName     = "/api.user.service.v1.User/RegisterWithCredentials"
	User_LoginWithCredentials_FullMethodName        = "/api.user.service.v1.User/LoginWithCredentials"
	User_RegisterOrLoginWithOAuth_FullMethodName    = "/api.user.service.v1.User/RegisterOrLoginWithOAuth"
	User_RegisterOrLoginWithTelegram_FullMethodName = "/api.user.service.v1.User/RegisterOrLoginWithTelegram"
	User_RefreshToken_FullMethodName                = "/api.user.service.v1.User/RefreshToken"
	User_GetUser_FullMethodName                     = "/api.user.service.v1.User/GetUser"
	User_Logout_FullMethodName                      = "/api.user.service.v1.User/Logout"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	RegisterWithCredentials(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	LoginWithCredentials(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	RegisterOrLoginWithOAuth(ctx context.Context, in *OAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	RegisterOrLoginWithTelegram(ctx context.Context, in *TelegramAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) RegisterWithCredentials(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, User_RegisterWithCredentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginWithCredentials(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, User_LoginWithCredentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RegisterOrLoginWithOAuth(ctx context.Context, in *OAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, User_RegisterOrLoginWithOAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RegisterOrLoginWithTelegram(ctx context.Context, in *TelegramAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, User_RegisterOrLoginWithTelegram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, User_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, User_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, User_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility.
type UserServer interface {
	RegisterWithCredentials(context.Context, *RegisterRequest) (*AuthResponse, error)
	LoginWithCredentials(context.Context, *LoginRequest) (*AuthResponse, error)
	RegisterOrLoginWithOAuth(context.Context, *OAuthRequest) (*AuthResponse, error)
	RegisterOrLoginWithTelegram(context.Context, *TelegramAuthRequest) (*AuthResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*AuthResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServer struct{}

func (UnimplementedUserServer) RegisterWithCredentials(context.Context, *RegisterRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWithCredentials not implemented")
}
func (UnimplementedUserServer) LoginWithCredentials(context.Context, *LoginRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithCredentials not implemented")
}
func (UnimplementedUserServer) RegisterOrLoginWithOAuth(context.Context, *OAuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterOrLoginWithOAuth not implemented")
}
func (UnimplementedUserServer) RegisterOrLoginWithTelegram(context.Context, *TelegramAuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterOrLoginWithTelegram not implemented")
}
func (UnimplementedUserServer) RefreshToken(context.Context, *RefreshTokenRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}
func (UnimplementedUserServer) testEmbeddedByValue()              {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	// If the following call pancis, it indicates UnimplementedUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_RegisterWithCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterWithCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RegisterWithCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterWithCredentials(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginWithCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginWithCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_LoginWithCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginWithCredentials(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RegisterOrLoginWithOAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterOrLoginWithOAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RegisterOrLoginWithOAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterOrLoginWithOAuth(ctx, req.(*OAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RegisterOrLoginWithTelegram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelegramAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterOrLoginWithTelegram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RegisterOrLoginWithTelegram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterOrLoginWithTelegram(ctx, req.(*TelegramAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.service.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWithCredentials",
			Handler:    _User_RegisterWithCredentials_Handler,
		},
		{
			MethodName: "LoginWithCredentials",
			Handler:    _User_LoginWithCredentials_Handler,
		},
		{
			MethodName: "RegisterOrLoginWithOAuth",
			Handler:    _User_RegisterOrLoginWithOAuth_Handler,
		},
		{
			MethodName: "RegisterOrLoginWithTelegram",
			Handler:    _User_RegisterOrLoginWithTelegram_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _User_RefreshToken_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _User_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/service/v1/user.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user/v1/user.proto

package user_v1

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *UserServiceAddUserRequest, opts ...grpc.CallOption) (*UserServiceAddUserResponse, error)
	GetUsers(ctx context.Context, in *UserServiceGetUsersRequest, opts ...grpc.CallOption) (*UserServiceGetUsersResponse, error)
	GetUser(ctx context.Context, in *UserServiceGetUserRequest, opts ...grpc.CallOption) (*UserServiceGetUserResponse, error)
	GetUsersStream(ctx context.Context, in *UserServiceGetUsersStreamRequest, opts ...grpc.CallOption) (UserService_GetUsersStreamClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *UserServiceAddUserRequest, opts ...grpc.CallOption) (*UserServiceAddUserResponse, error) {
	out := new(UserServiceAddUserResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *UserServiceGetUsersRequest, opts ...grpc.CallOption) (*UserServiceGetUsersResponse, error) {
	out := new(UserServiceGetUsersResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserServiceGetUserRequest, opts ...grpc.CallOption) (*UserServiceGetUserResponse, error) {
	out := new(UserServiceGetUserResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsersStream(ctx context.Context, in *UserServiceGetUsersStreamRequest, opts ...grpc.CallOption) (UserService_GetUsersStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/user.v1.UserService/GetUsersStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUsersStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUsersStreamClient interface {
	Recv() (*UserServiceGetUsersStreamResponse, error)
	grpc.ClientStream
}

type userServiceGetUsersStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUsersStreamClient) Recv() (*UserServiceGetUsersStreamResponse, error) {
	m := new(UserServiceGetUsersStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	AddUser(context.Context, *UserServiceAddUserRequest) (*UserServiceAddUserResponse, error)
	GetUsers(context.Context, *UserServiceGetUsersRequest) (*UserServiceGetUsersResponse, error)
	GetUser(context.Context, *UserServiceGetUserRequest) (*UserServiceGetUserResponse, error)
	GetUsersStream(*UserServiceGetUsersStreamRequest, UserService_GetUsersStreamServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) AddUser(context.Context, *UserServiceAddUserRequest) (*UserServiceAddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServiceServer) GetUsers(context.Context, *UserServiceGetUsersRequest) (*UserServiceGetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *UserServiceGetUserRequest) (*UserServiceGetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetUsersStream(*UserServiceGetUsersStreamRequest, UserService_GetUsersStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsersStream not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceAddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*UserServiceAddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceGetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*UserServiceGetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceGetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserServiceGetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsersStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserServiceGetUsersStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUsersStream(m, &userServiceGetUsersStreamServer{stream})
}

type UserService_GetUsersStreamServer interface {
	Send(*UserServiceGetUsersStreamResponse) error
	grpc.ServerStream
}

type userServiceGetUsersStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUsersStreamServer) Send(m *UserServiceGetUsersStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsersStream",
			Handler:       _UserService_GetUsersStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user/v1/user.proto",
}

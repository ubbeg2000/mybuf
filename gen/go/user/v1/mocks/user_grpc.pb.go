// Code generated by MockGen. DO NOT EDIT.
// Source: gen/go/user/v1/user_grpc.pb.go

// Package mock_user_v1 is a generated GoMock package.
package mock_user_v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user_v1 "github.com/ubbeg2000/mybuf/gen/go/user/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserServiceClient) AddUser(ctx context.Context, in *user_v1.UserServiceAddUserRequest, opts ...grpc.CallOption) (*user_v1.UserServiceAddUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUser", varargs...)
	ret0, _ := ret[0].(*user_v1.UserServiceAddUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServiceClientMockRecorder) AddUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserServiceClient)(nil).AddUser), varargs...)
}

// GetUser mocks base method.
func (m *MockUserServiceClient) GetUser(ctx context.Context, in *user_v1.UserServiceGetUserRequest, opts ...grpc.CallOption) (*user_v1.UserServiceGetUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*user_v1.UserServiceGetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceClient)(nil).GetUser), varargs...)
}

// GetUsers mocks base method.
func (m *MockUserServiceClient) GetUsers(ctx context.Context, in *user_v1.UserServiceGetUsersRequest, opts ...grpc.CallOption) (*user_v1.UserServiceGetUsersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsers", varargs...)
	ret0, _ := ret[0].(*user_v1.UserServiceGetUsersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceClientMockRecorder) GetUsers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserServiceClient)(nil).GetUsers), varargs...)
}

// GetUsersStream mocks base method.
func (m *MockUserServiceClient) GetUsersStream(ctx context.Context, in *user_v1.UserServiceGetUsersStreamRequest, opts ...grpc.CallOption) (user_v1.UserService_GetUsersStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersStream", varargs...)
	ret0, _ := ret[0].(user_v1.UserService_GetUsersStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersStream indicates an expected call of GetUsersStream.
func (mr *MockUserServiceClientMockRecorder) GetUsersStream(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersStream", reflect.TypeOf((*MockUserServiceClient)(nil).GetUsersStream), varargs...)
}

// MockUserService_GetUsersStreamClient is a mock of UserService_GetUsersStreamClient interface.
type MockUserService_GetUsersStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_GetUsersStreamClientMockRecorder
}

// MockUserService_GetUsersStreamClientMockRecorder is the mock recorder for MockUserService_GetUsersStreamClient.
type MockUserService_GetUsersStreamClientMockRecorder struct {
	mock *MockUserService_GetUsersStreamClient
}

// NewMockUserService_GetUsersStreamClient creates a new mock instance.
func NewMockUserService_GetUsersStreamClient(ctrl *gomock.Controller) *MockUserService_GetUsersStreamClient {
	mock := &MockUserService_GetUsersStreamClient{ctrl: ctrl}
	mock.recorder = &MockUserService_GetUsersStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService_GetUsersStreamClient) EXPECT() *MockUserService_GetUsersStreamClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockUserService_GetUsersStreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockUserService_GetUsersStreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockUserService_GetUsersStreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockUserService_GetUsersStreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUserService_GetUsersStreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserService_GetUsersStreamClient)(nil).Context))
}

// Header mocks base method.
func (m *MockUserService_GetUsersStreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockUserService_GetUsersStreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockUserService_GetUsersStreamClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockUserService_GetUsersStreamClient) Recv() (*user_v1.UserServiceGetUsersStreamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*user_v1.UserServiceGetUsersStreamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockUserService_GetUsersStreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockUserService_GetUsersStreamClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockUserService_GetUsersStreamClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUserService_GetUsersStreamClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_GetUsersStreamClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockUserService_GetUsersStreamClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUserService_GetUsersStreamClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_GetUsersStreamClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockUserService_GetUsersStreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockUserService_GetUsersStreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockUserService_GetUsersStreamClient)(nil).Trailer))
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserServiceServer) AddUser(arg0 context.Context, arg1 *user_v1.UserServiceAddUserRequest) (*user_v1.UserServiceAddUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(*user_v1.UserServiceAddUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServiceServerMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserServiceServer)(nil).AddUser), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockUserServiceServer) GetUser(arg0 context.Context, arg1 *user_v1.UserServiceGetUserRequest) (*user_v1.UserServiceGetUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*user_v1.UserServiceGetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceServer)(nil).GetUser), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockUserServiceServer) GetUsers(arg0 context.Context, arg1 *user_v1.UserServiceGetUsersRequest) (*user_v1.UserServiceGetUsersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].(*user_v1.UserServiceGetUsersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceServerMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserServiceServer)(nil).GetUsers), arg0, arg1)
}

// GetUsersStream mocks base method.
func (m *MockUserServiceServer) GetUsersStream(arg0 *user_v1.UserServiceGetUsersStreamRequest, arg1 user_v1.UserService_GetUsersStreamServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUsersStream indicates an expected call of GetUsersStream.
func (mr *MockUserServiceServerMockRecorder) GetUsersStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersStream", reflect.TypeOf((*MockUserServiceServer)(nil).GetUsersStream), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUserService_GetUsersStreamServer is a mock of UserService_GetUsersStreamServer interface.
type MockUserService_GetUsersStreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_GetUsersStreamServerMockRecorder
}

// MockUserService_GetUsersStreamServerMockRecorder is the mock recorder for MockUserService_GetUsersStreamServer.
type MockUserService_GetUsersStreamServerMockRecorder struct {
	mock *MockUserService_GetUsersStreamServer
}

// NewMockUserService_GetUsersStreamServer creates a new mock instance.
func NewMockUserService_GetUsersStreamServer(ctrl *gomock.Controller) *MockUserService_GetUsersStreamServer {
	mock := &MockUserService_GetUsersStreamServer{ctrl: ctrl}
	mock.recorder = &MockUserService_GetUsersStreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService_GetUsersStreamServer) EXPECT() *MockUserService_GetUsersStreamServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockUserService_GetUsersStreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUserService_GetUsersStreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserService_GetUsersStreamServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockUserService_GetUsersStreamServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUserService_GetUsersStreamServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_GetUsersStreamServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockUserService_GetUsersStreamServer) Send(arg0 *user_v1.UserServiceGetUsersStreamResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockUserService_GetUsersStreamServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockUserService_GetUsersStreamServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockUserService_GetUsersStreamServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockUserService_GetUsersStreamServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockUserService_GetUsersStreamServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockUserService_GetUsersStreamServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUserService_GetUsersStreamServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_GetUsersStreamServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockUserService_GetUsersStreamServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockUserService_GetUsersStreamServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockUserService_GetUsersStreamServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockUserService_GetUsersStreamServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockUserService_GetUsersStreamServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockUserService_GetUsersStreamServer)(nil).SetTrailer), arg0)
}

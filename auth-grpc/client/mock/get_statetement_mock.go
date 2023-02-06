// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Edbeer/payment-proto/auth-grpc/proto (interfaces: AuthService_GetStatementClient,AuthService_GetStatementServer)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	reflect "reflect"

	authpb "github.com/Edbeer/payment-proto/auth-grpc/proto"
	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockAuthService_GetStatementClient is a mock of AuthService_GetStatementClient interface.
type MockAuthService_GetStatementClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthService_GetStatementClientMockRecorder
}

// MockAuthService_GetStatementClientMockRecorder is the mock recorder for MockAuthService_GetStatementClient.
type MockAuthService_GetStatementClientMockRecorder struct {
	mock *MockAuthService_GetStatementClient
}

// NewMockAuthService_GetStatementClient creates a new mock instance.
func NewMockAuthService_GetStatementClient(ctrl *gomock.Controller) *MockAuthService_GetStatementClient {
	mock := &MockAuthService_GetStatementClient{ctrl: ctrl}
	mock.recorder = &MockAuthService_GetStatementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService_GetStatementClient) EXPECT() *MockAuthService_GetStatementClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockAuthService_GetStatementClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockAuthService_GetStatementClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockAuthService_GetStatementClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockAuthService_GetStatementClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAuthService_GetStatementClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAuthService_GetStatementClient)(nil).Context))
}

// Header mocks base method.
func (m *MockAuthService_GetStatementClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockAuthService_GetStatementClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockAuthService_GetStatementClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockAuthService_GetStatementClient) Recv() (*authpb.Statement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*authpb.Statement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockAuthService_GetStatementClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAuthService_GetStatementClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockAuthService_GetStatementClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAuthService_GetStatementClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAuthService_GetStatementClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockAuthService_GetStatementClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAuthService_GetStatementClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAuthService_GetStatementClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockAuthService_GetStatementClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockAuthService_GetStatementClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockAuthService_GetStatementClient)(nil).Trailer))
}

// MockAuthService_GetStatementServer is a mock of AuthService_GetStatementServer interface.
type MockAuthService_GetStatementServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthService_GetStatementServerMockRecorder
}

// MockAuthService_GetStatementServerMockRecorder is the mock recorder for MockAuthService_GetStatementServer.
type MockAuthService_GetStatementServerMockRecorder struct {
	mock *MockAuthService_GetStatementServer
}

// NewMockAuthService_GetStatementServer creates a new mock instance.
func NewMockAuthService_GetStatementServer(ctrl *gomock.Controller) *MockAuthService_GetStatementServer {
	mock := &MockAuthService_GetStatementServer{ctrl: ctrl}
	mock.recorder = &MockAuthService_GetStatementServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService_GetStatementServer) EXPECT() *MockAuthService_GetStatementServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockAuthService_GetStatementServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAuthService_GetStatementServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAuthService_GetStatementServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockAuthService_GetStatementServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAuthService_GetStatementServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAuthService_GetStatementServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockAuthService_GetStatementServer) Send(arg0 *authpb.Statement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAuthService_GetStatementServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAuthService_GetStatementServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockAuthService_GetStatementServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockAuthService_GetStatementServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockAuthService_GetStatementServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockAuthService_GetStatementServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAuthService_GetStatementServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAuthService_GetStatementServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockAuthService_GetStatementServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockAuthService_GetStatementServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockAuthService_GetStatementServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockAuthService_GetStatementServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockAuthService_GetStatementServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockAuthService_GetStatementServer)(nil).SetTrailer), arg0)
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: auth.proto

package authpb

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*AccountWithTokens, error)
	SignIn(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AccountWithTokens, error)
	SignOut(ctx context.Context, in *QuitRequest, opts ...grpc.CallOption) (*QuitResponse, error)
	RefreshTokens(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*Tokens, error)
	GetAccount(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (AuthService_GetAccountClient, error)
	UpdateAccount(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Account, error)
	DeleteAccount(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DepositAccount(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
	// for payment
	GetAccountByID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*Account, error)
	GetStatement(ctx context.Context, in *StatementGet, opts ...grpc.CallOption) (AuthService_GetStatementClient, error)
	CreateStatement(ctx context.Context, opts ...grpc.CallOption) (AuthService_CreateStatementClient, error)
	UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*Account, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateAccount(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*AccountWithTokens, error) {
	out := new(AccountWithTokens)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignIn(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AccountWithTokens, error) {
	out := new(AccountWithTokens)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignOut(ctx context.Context, in *QuitRequest, opts ...grpc.CallOption) (*QuitResponse, error) {
	out := new(QuitResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshTokens(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/auth.AuthService/RefreshTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAccount(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (AuthService_GetAccountClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuthService_ServiceDesc.Streams[0], "/auth.AuthService/GetAccount", opts...)
	if err != nil {
		return nil, err
	}
	x := &authServiceGetAccountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuthService_GetAccountClient interface {
	Recv() (*Account, error)
	grpc.ClientStream
}

type authServiceGetAccountClient struct {
	grpc.ClientStream
}

func (x *authServiceGetAccountClient) Recv() (*Account, error) {
	m := new(Account)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authServiceClient) UpdateAccount(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteAccount(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DepositAccount(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DepositAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAccountByID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetAccountByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetStatement(ctx context.Context, in *StatementGet, opts ...grpc.CallOption) (AuthService_GetStatementClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuthService_ServiceDesc.Streams[1], "/auth.AuthService/GetStatement", opts...)
	if err != nil {
		return nil, err
	}
	x := &authServiceGetStatementClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuthService_GetStatementClient interface {
	Recv() (*Statement, error)
	grpc.ClientStream
}

type authServiceGetStatementClient struct {
	grpc.ClientStream
}

func (x *authServiceGetStatementClient) Recv() (*Statement, error) {
	m := new(Statement)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authServiceClient) CreateStatement(ctx context.Context, opts ...grpc.CallOption) (AuthService_CreateStatementClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuthService_ServiceDesc.Streams[2], "/auth.AuthService/CreateStatement", opts...)
	if err != nil {
		return nil, err
	}
	x := &authServiceCreateStatementClient{stream}
	return x, nil
}

type AuthService_CreateStatementClient interface {
	Send(*StatementRequest) error
	Recv() (*StatementResponse, error)
	grpc.ClientStream
}

type authServiceCreateStatementClient struct {
	grpc.ClientStream
}

func (x *authServiceCreateStatementClient) Send(m *StatementRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *authServiceCreateStatementClient) Recv() (*StatementResponse, error) {
	m := new(StatementResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authServiceClient) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	CreateAccount(context.Context, *CreateRequest) (*AccountWithTokens, error)
	SignIn(context.Context, *LoginRequest) (*AccountWithTokens, error)
	SignOut(context.Context, *QuitRequest) (*QuitResponse, error)
	RefreshTokens(context.Context, *RefreshRequest) (*Tokens, error)
	GetAccount(*GetRequest, AuthService_GetAccountServer) error
	UpdateAccount(context.Context, *UpdateRequest) (*Account, error)
	DeleteAccount(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DepositAccount(context.Context, *DepositRequest) (*DepositResponse, error)
	// for payment
	GetAccountByID(context.Context, *GetIDRequest) (*Account, error)
	GetStatement(*StatementGet, AuthService_GetStatementServer) error
	CreateStatement(AuthService_CreateStatementServer) error
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*Account, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) CreateAccount(context.Context, *CreateRequest) (*AccountWithTokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAuthServiceServer) SignIn(context.Context, *LoginRequest) (*AccountWithTokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServiceServer) SignOut(context.Context, *QuitRequest) (*QuitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedAuthServiceServer) RefreshTokens(context.Context, *RefreshRequest) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTokens not implemented")
}
func (UnimplementedAuthServiceServer) GetAccount(*GetRequest, AuthService_GetAccountServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAuthServiceServer) UpdateAccount(context.Context, *UpdateRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAuthServiceServer) DeleteAccount(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAuthServiceServer) DepositAccount(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositAccount not implemented")
}
func (UnimplementedAuthServiceServer) GetAccountByID(context.Context, *GetIDRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByID not implemented")
}
func (UnimplementedAuthServiceServer) GetStatement(*StatementGet, AuthService_GetStatementServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatement not implemented")
}
func (UnimplementedAuthServiceServer) CreateStatement(AuthService_CreateStatementServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateStatement not implemented")
}
func (UnimplementedAuthServiceServer) UpdateBalance(context.Context, *UpdateBalanceRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateAccount(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignIn(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignOut(ctx, req.(*QuitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/RefreshTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshTokens(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAccount_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthServiceServer).GetAccount(m, &authServiceGetAccountServer{stream})
}

type AuthService_GetAccountServer interface {
	Send(*Account) error
	grpc.ServerStream
}

type authServiceGetAccountServer struct {
	grpc.ServerStream
}

func (x *authServiceGetAccountServer) Send(m *Account) error {
	return x.ServerStream.SendMsg(m)
}

func _AuthService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateAccount(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteAccount(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DepositAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DepositAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DepositAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DepositAccount(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAccountByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAccountByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetAccountByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAccountByID(ctx, req.(*GetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetStatement_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatementGet)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthServiceServer).GetStatement(m, &authServiceGetStatementServer{stream})
}

type AuthService_GetStatementServer interface {
	Send(*Statement) error
	grpc.ServerStream
}

type authServiceGetStatementServer struct {
	grpc.ServerStream
}

func (x *authServiceGetStatementServer) Send(m *Statement) error {
	return x.ServerStream.SendMsg(m)
}

func _AuthService_CreateStatement_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuthServiceServer).CreateStatement(&authServiceCreateStatementServer{stream})
}

type AuthService_CreateStatementServer interface {
	Send(*StatementResponse) error
	Recv() (*StatementRequest, error)
	grpc.ServerStream
}

type authServiceCreateStatementServer struct {
	grpc.ServerStream
}

func (x *authServiceCreateStatementServer) Send(m *StatementResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *authServiceCreateStatementServer) Recv() (*StatementRequest, error) {
	m := new(StatementRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AuthService_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateBalance(ctx, req.(*UpdateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AuthService_CreateAccount_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthService_SignOut_Handler,
		},
		{
			MethodName: "RefreshTokens",
			Handler:    _AuthService_RefreshTokens_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AuthService_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AuthService_DeleteAccount_Handler,
		},
		{
			MethodName: "DepositAccount",
			Handler:    _AuthService_DepositAccount_Handler,
		},
		{
			MethodName: "GetAccountByID",
			Handler:    _AuthService_GetAccountByID_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _AuthService_UpdateBalance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAccount",
			Handler:       _AuthService_GetAccount_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStatement",
			Handler:       _AuthService_GetStatement_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateStatement",
			Handler:       _AuthService_CreateStatement_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "auth.proto",
}

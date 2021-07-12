// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stub

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/service.ChatService/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatClient{stream}
	return x, nil
}

type ChatService_ChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessageFromServer, error)
	grpc.ClientStream
}

type chatServiceChatClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChatClient) Recv() (*ChatMessageFromServer, error) {
	m := new(ChatMessageFromServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Chat(ChatService_ChatServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Chat(ChatService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Chat(&chatServiceChatServer{stream})
}

type ChatService_ChatServer interface {
	Send(*ChatMessageFromServer) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatServiceChatServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatServer) Send(m *ChatMessageFromServer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _ChatService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/service/messageService.proto",
}

// AuthenticateUserClient is the client API for AuthenticateUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticateUserClient interface {
	Register(ctx context.Context, in *RegisterUser, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Login(ctx context.Context, in *LoginUser, opts ...grpc.CallOption) (*Token, error)
}

type authenticateUserClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticateUserClient(cc grpc.ClientConnInterface) AuthenticateUserClient {
	return &authenticateUserClient{cc}
}

func (c *authenticateUserClient) Register(ctx context.Context, in *RegisterUser, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.AuthenticateUser/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateUserClient) Login(ctx context.Context, in *LoginUser, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/service.AuthenticateUser/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateUserServer is the server API for AuthenticateUser service.
// All implementations must embed UnimplementedAuthenticateUserServer
// for forward compatibility
type AuthenticateUserServer interface {
	Register(context.Context, *RegisterUser) (*emptypb.Empty, error)
	Login(context.Context, *LoginUser) (*Token, error)
	mustEmbedUnimplementedAuthenticateUserServer()
}

// UnimplementedAuthenticateUserServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticateUserServer struct {
}

func (UnimplementedAuthenticateUserServer) Register(context.Context, *RegisterUser) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthenticateUserServer) Login(context.Context, *LoginUser) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticateUserServer) mustEmbedUnimplementedAuthenticateUserServer() {}

// UnsafeAuthenticateUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticateUserServer will
// result in compilation errors.
type UnsafeAuthenticateUserServer interface {
	mustEmbedUnimplementedAuthenticateUserServer()
}

func RegisterAuthenticateUserServer(s grpc.ServiceRegistrar, srv AuthenticateUserServer) {
	s.RegisterService(&AuthenticateUser_ServiceDesc, srv)
}

func _AuthenticateUser_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateUserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AuthenticateUser/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateUserServer).Register(ctx, req.(*RegisterUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateUser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateUserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AuthenticateUser/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateUserServer).Login(ctx, req.(*LoginUser))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticateUser_ServiceDesc is the grpc.ServiceDesc for AuthenticateUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticateUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.AuthenticateUser",
	HandlerType: (*AuthenticateUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthenticateUser_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthenticateUser_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/service/messageService.proto",
}

// UpdateUserClient is the client API for UpdateUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateUserClient interface {
	UpdateName(ctx context.Context, in *Edit, opts ...grpc.CallOption) (*RegisterUser, error)
	AddFriend(ctx context.Context, in *FriendList, opts ...grpc.CallOption) (*FriendList, error)
}

type updateUserClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateUserClient(cc grpc.ClientConnInterface) UpdateUserClient {
	return &updateUserClient{cc}
}

func (c *updateUserClient) UpdateName(ctx context.Context, in *Edit, opts ...grpc.CallOption) (*RegisterUser, error) {
	out := new(RegisterUser)
	err := c.cc.Invoke(ctx, "/service.UpdateUser/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateUserClient) AddFriend(ctx context.Context, in *FriendList, opts ...grpc.CallOption) (*FriendList, error) {
	out := new(FriendList)
	err := c.cc.Invoke(ctx, "/service.UpdateUser/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateUserServer is the server API for UpdateUser service.
// All implementations must embed UnimplementedUpdateUserServer
// for forward compatibility
type UpdateUserServer interface {
	UpdateName(context.Context, *Edit) (*RegisterUser, error)
	AddFriend(context.Context, *FriendList) (*FriendList, error)
	mustEmbedUnimplementedUpdateUserServer()
}

// UnimplementedUpdateUserServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateUserServer struct {
}

func (UnimplementedUpdateUserServer) UpdateName(context.Context, *Edit) (*RegisterUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedUpdateUserServer) AddFriend(context.Context, *FriendList) (*FriendList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedUpdateUserServer) mustEmbedUnimplementedUpdateUserServer() {}

// UnsafeUpdateUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateUserServer will
// result in compilation errors.
type UnsafeUpdateUserServer interface {
	mustEmbedUnimplementedUpdateUserServer()
}

func RegisterUpdateUserServer(s grpc.ServiceRegistrar, srv UpdateUserServer) {
	s.RegisterService(&UpdateUser_ServiceDesc, srv)
}

func _UpdateUser_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Edit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateUserServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UpdateUser/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateUserServer).UpdateName(ctx, req.(*Edit))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpdateUser_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateUserServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UpdateUser/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateUserServer).AddFriend(ctx, req.(*FriendList))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateUser_ServiceDesc is the grpc.ServiceDesc for UpdateUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.UpdateUser",
	HandlerType: (*UpdateUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateName",
			Handler:    _UpdateUser_UpdateName_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _UpdateUser_AddFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/service/messageService.proto",
}
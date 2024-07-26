// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gophkeeper

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

// GophkeeperClient is the client API for Gophkeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophkeeperClient interface {
	RegisterUser(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AuthenticateUser(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*AuthorizationToken, error)
	SaveLoginPassword(ctx context.Context, in *LoginPasswordData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SaveCard(ctx context.Context, in *CardData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SaveText(ctx context.Context, in *TextData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SaveBinary(ctx context.Context, in *BinaryData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetListData(ctx context.Context, in *DataGetterRequest, opts ...grpc.CallOption) (*DataListResponse, error)
	GetLoginPasswordObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*LoginPasswordData, error)
	GetCardObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*CardData, error)
	GetTextObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*TextData, error)
	GetBinaryObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*BinaryData, error)
	DeleteObject(ctx context.Context, in *DataGetterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gophkeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophkeeperClient(cc grpc.ClientConnInterface) GophkeeperClient {
	return &gophkeeperClient{cc}
}

func (c *gophkeeperClient) RegisterUser(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) AuthenticateUser(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*AuthorizationToken, error) {
	out := new(AuthorizationToken)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/AuthenticateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) SaveLoginPassword(ctx context.Context, in *LoginPasswordData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/SaveLoginPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) SaveCard(ctx context.Context, in *CardData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/SaveCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) SaveText(ctx context.Context, in *TextData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/SaveText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) SaveBinary(ctx context.Context, in *BinaryData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/SaveBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetListData(ctx context.Context, in *DataGetterRequest, opts ...grpc.CallOption) (*DataListResponse, error) {
	out := new(DataListResponse)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/GetListData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetLoginPasswordObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*LoginPasswordData, error) {
	out := new(LoginPasswordData)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/GetLoginPasswordObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetCardObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*CardData, error) {
	out := new(CardData)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/GetCardObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetTextObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*TextData, error) {
	out := new(TextData)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/GetTextObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetBinaryObject(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*BinaryData, error) {
	out := new(BinaryData)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/GetBinaryObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) DeleteObject(ctx context.Context, in *DataGetterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/shortener.Gophkeeper/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophkeeperServer is the server API for Gophkeeper service.
// All implementations must embed UnimplementedGophkeeperServer
// for forward compatibility
type GophkeeperServer interface {
	RegisterUser(context.Context, *UserCredentials) (*emptypb.Empty, error)
	AuthenticateUser(context.Context, *UserCredentials) (*AuthorizationToken, error)
	SaveLoginPassword(context.Context, *LoginPasswordData) (*emptypb.Empty, error)
	SaveCard(context.Context, *CardData) (*emptypb.Empty, error)
	SaveText(context.Context, *TextData) (*emptypb.Empty, error)
	SaveBinary(context.Context, *BinaryData) (*emptypb.Empty, error)
	GetListData(context.Context, *DataGetterRequest) (*DataListResponse, error)
	GetLoginPasswordObject(context.Context, *DataRequest) (*LoginPasswordData, error)
	GetCardObject(context.Context, *DataRequest) (*CardData, error)
	GetTextObject(context.Context, *DataRequest) (*TextData, error)
	GetBinaryObject(context.Context, *DataRequest) (*BinaryData, error)
	DeleteObject(context.Context, *DataGetterRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGophkeeperServer()
}

// UnimplementedGophkeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophkeeperServer struct {
}

func (UnimplementedGophkeeperServer) RegisterUser(context.Context, *UserCredentials) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedGophkeeperServer) AuthenticateUser(context.Context, *UserCredentials) (*AuthorizationToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUser not implemented")
}
func (UnimplementedGophkeeperServer) SaveLoginPassword(context.Context, *LoginPasswordData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLoginPassword not implemented")
}
func (UnimplementedGophkeeperServer) SaveCard(context.Context, *CardData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCard not implemented")
}
func (UnimplementedGophkeeperServer) SaveText(context.Context, *TextData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveText not implemented")
}
func (UnimplementedGophkeeperServer) SaveBinary(context.Context, *BinaryData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBinary not implemented")
}
func (UnimplementedGophkeeperServer) GetListData(context.Context, *DataGetterRequest) (*DataListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListData not implemented")
}
func (UnimplementedGophkeeperServer) GetLoginPasswordObject(context.Context, *DataRequest) (*LoginPasswordData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginPasswordObject not implemented")
}
func (UnimplementedGophkeeperServer) GetCardObject(context.Context, *DataRequest) (*CardData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardObject not implemented")
}
func (UnimplementedGophkeeperServer) GetTextObject(context.Context, *DataRequest) (*TextData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTextObject not implemented")
}
func (UnimplementedGophkeeperServer) GetBinaryObject(context.Context, *DataRequest) (*BinaryData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBinaryObject not implemented")
}
func (UnimplementedGophkeeperServer) DeleteObject(context.Context, *DataGetterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedGophkeeperServer) mustEmbedUnimplementedGophkeeperServer() {}

// UnsafeGophkeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophkeeperServer will
// result in compilation errors.
type UnsafeGophkeeperServer interface {
	mustEmbedUnimplementedGophkeeperServer()
}

func RegisterGophkeeperServer(s grpc.ServiceRegistrar, srv GophkeeperServer) {
	s.RegisterService(&Gophkeeper_ServiceDesc, srv)
}

func _Gophkeeper_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).RegisterUser(ctx, req.(*UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_AuthenticateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).AuthenticateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/AuthenticateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).AuthenticateUser(ctx, req.(*UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_SaveLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPasswordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).SaveLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/SaveLoginPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).SaveLoginPassword(ctx, req.(*LoginPasswordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_SaveCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).SaveCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/SaveCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).SaveCard(ctx, req.(*CardData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_SaveText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).SaveText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/SaveText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).SaveText(ctx, req.(*TextData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_SaveBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).SaveBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/SaveBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).SaveBinary(ctx, req.(*BinaryData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetListData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataGetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetListData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/GetListData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetListData(ctx, req.(*DataGetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetLoginPasswordObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetLoginPasswordObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/GetLoginPasswordObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetLoginPasswordObject(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetCardObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetCardObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/GetCardObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetCardObject(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetTextObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetTextObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/GetTextObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetTextObject(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetBinaryObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetBinaryObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/GetBinaryObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetBinaryObject(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataGetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Gophkeeper/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).DeleteObject(ctx, req.(*DataGetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gophkeeper_ServiceDesc is the grpc.ServiceDesc for Gophkeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gophkeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener.Gophkeeper",
	HandlerType: (*GophkeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _Gophkeeper_RegisterUser_Handler,
		},
		{
			MethodName: "AuthenticateUser",
			Handler:    _Gophkeeper_AuthenticateUser_Handler,
		},
		{
			MethodName: "SaveLoginPassword",
			Handler:    _Gophkeeper_SaveLoginPassword_Handler,
		},
		{
			MethodName: "SaveCard",
			Handler:    _Gophkeeper_SaveCard_Handler,
		},
		{
			MethodName: "SaveText",
			Handler:    _Gophkeeper_SaveText_Handler,
		},
		{
			MethodName: "SaveBinary",
			Handler:    _Gophkeeper_SaveBinary_Handler,
		},
		{
			MethodName: "GetListData",
			Handler:    _Gophkeeper_GetListData_Handler,
		},
		{
			MethodName: "GetLoginPasswordObject",
			Handler:    _Gophkeeper_GetLoginPasswordObject_Handler,
		},
		{
			MethodName: "GetCardObject",
			Handler:    _Gophkeeper_GetCardObject_Handler,
		},
		{
			MethodName: "GetTextObject",
			Handler:    _Gophkeeper_GetTextObject_Handler,
		},
		{
			MethodName: "GetBinaryObject",
			Handler:    _Gophkeeper_GetBinaryObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Gophkeeper_DeleteObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gophkeeper.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: rpcMsg.proto

package rpcMsg

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

// ControlMsgServiceClient is the client API for ControlMsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlMsgServiceClient interface {
	HandleHelloMsg(ctx context.Context, in *RequestHello, opts ...grpc.CallOption) (*ResponseWelcome, error)
	HandleSetStationMsg(ctx context.Context, in *RequestSetStation, opts ...grpc.CallOption) (*ResponseAnnounce, error)
}

type controlMsgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControlMsgServiceClient(cc grpc.ClientConnInterface) ControlMsgServiceClient {
	return &controlMsgServiceClient{cc}
}

func (c *controlMsgServiceClient) HandleHelloMsg(ctx context.Context, in *RequestHello, opts ...grpc.CallOption) (*ResponseWelcome, error) {
	out := new(ResponseWelcome)
	err := c.cc.Invoke(ctx, "/rpcMsg.ControlMsgService/HandleHelloMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlMsgServiceClient) HandleSetStationMsg(ctx context.Context, in *RequestSetStation, opts ...grpc.CallOption) (*ResponseAnnounce, error) {
	out := new(ResponseAnnounce)
	err := c.cc.Invoke(ctx, "/rpcMsg.ControlMsgService/HandleSetStationMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlMsgServiceServer is the server API for ControlMsgService service.
// All implementations must embed UnimplementedControlMsgServiceServer
// for forward compatibility
type ControlMsgServiceServer interface {
	HandleHelloMsg(context.Context, *RequestHello) (*ResponseWelcome, error)
	HandleSetStationMsg(context.Context, *RequestSetStation) (*ResponseAnnounce, error)
	mustEmbedUnimplementedControlMsgServiceServer()
}

// UnimplementedControlMsgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedControlMsgServiceServer struct {
}

func (UnimplementedControlMsgServiceServer) HandleHelloMsg(context.Context, *RequestHello) (*ResponseWelcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleHelloMsg not implemented")
}
func (UnimplementedControlMsgServiceServer) HandleSetStationMsg(context.Context, *RequestSetStation) (*ResponseAnnounce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSetStationMsg not implemented")
}
func (UnimplementedControlMsgServiceServer) mustEmbedUnimplementedControlMsgServiceServer() {}

// UnsafeControlMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlMsgServiceServer will
// result in compilation errors.
type UnsafeControlMsgServiceServer interface {
	mustEmbedUnimplementedControlMsgServiceServer()
}

func RegisterControlMsgServiceServer(s grpc.ServiceRegistrar, srv ControlMsgServiceServer) {
	s.RegisterService(&ControlMsgService_ServiceDesc, srv)
}

func _ControlMsgService_HandleHelloMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestHello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlMsgServiceServer).HandleHelloMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcMsg.ControlMsgService/HandleHelloMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlMsgServiceServer).HandleHelloMsg(ctx, req.(*RequestHello))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlMsgService_HandleSetStationMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSetStation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlMsgServiceServer).HandleSetStationMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcMsg.ControlMsgService/HandleSetStationMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlMsgServiceServer).HandleSetStationMsg(ctx, req.(*RequestSetStation))
	}
	return interceptor(ctx, in, info, handler)
}

// ControlMsgService_ServiceDesc is the grpc.ServiceDesc for ControlMsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControlMsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcMsg.ControlMsgService",
	HandlerType: (*ControlMsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleHelloMsg",
			Handler:    _ControlMsgService_HandleHelloMsg_Handler,
		},
		{
			MethodName: "HandleSetStationMsg",
			Handler:    _ControlMsgService_HandleSetStationMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpcMsg.proto",
}

// ServerMsgServiceClient is the client API for ServerMsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerMsgServiceClient interface {
	HandleSendFileMsg(ctx context.Context, in *RequestSendFile, opts ...grpc.CallOption) (*ResponseSendFile, error)
	HandleShutdownMsg(ctx context.Context, in *RequestShutdown, opts ...grpc.CallOption) (*ResponseShutdown, error)
}

type serverMsgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerMsgServiceClient(cc grpc.ClientConnInterface) ServerMsgServiceClient {
	return &serverMsgServiceClient{cc}
}

func (c *serverMsgServiceClient) HandleSendFileMsg(ctx context.Context, in *RequestSendFile, opts ...grpc.CallOption) (*ResponseSendFile, error) {
	out := new(ResponseSendFile)
	err := c.cc.Invoke(ctx, "/rpcMsg.ServerMsgService/HandleSendFileMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverMsgServiceClient) HandleShutdownMsg(ctx context.Context, in *RequestShutdown, opts ...grpc.CallOption) (*ResponseShutdown, error) {
	out := new(ResponseShutdown)
	err := c.cc.Invoke(ctx, "/rpcMsg.ServerMsgService/HandleShutdownMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerMsgServiceServer is the server API for ServerMsgService service.
// All implementations must embed UnimplementedServerMsgServiceServer
// for forward compatibility
type ServerMsgServiceServer interface {
	HandleSendFileMsg(context.Context, *RequestSendFile) (*ResponseSendFile, error)
	HandleShutdownMsg(context.Context, *RequestShutdown) (*ResponseShutdown, error)
	mustEmbedUnimplementedServerMsgServiceServer()
}

// UnimplementedServerMsgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerMsgServiceServer struct {
}

func (UnimplementedServerMsgServiceServer) HandleSendFileMsg(context.Context, *RequestSendFile) (*ResponseSendFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSendFileMsg not implemented")
}
func (UnimplementedServerMsgServiceServer) HandleShutdownMsg(context.Context, *RequestShutdown) (*ResponseShutdown, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleShutdownMsg not implemented")
}
func (UnimplementedServerMsgServiceServer) mustEmbedUnimplementedServerMsgServiceServer() {}

// UnsafeServerMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerMsgServiceServer will
// result in compilation errors.
type UnsafeServerMsgServiceServer interface {
	mustEmbedUnimplementedServerMsgServiceServer()
}

func RegisterServerMsgServiceServer(s grpc.ServiceRegistrar, srv ServerMsgServiceServer) {
	s.RegisterService(&ServerMsgService_ServiceDesc, srv)
}

func _ServerMsgService_HandleSendFileMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSendFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerMsgServiceServer).HandleSendFileMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcMsg.ServerMsgService/HandleSendFileMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerMsgServiceServer).HandleSendFileMsg(ctx, req.(*RequestSendFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerMsgService_HandleShutdownMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestShutdown)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerMsgServiceServer).HandleShutdownMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcMsg.ServerMsgService/HandleShutdownMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerMsgServiceServer).HandleShutdownMsg(ctx, req.(*RequestShutdown))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerMsgService_ServiceDesc is the grpc.ServiceDesc for ServerMsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerMsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcMsg.ServerMsgService",
	HandlerType: (*ServerMsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleSendFileMsg",
			Handler:    _ServerMsgService_HandleSendFileMsg_Handler,
		},
		{
			MethodName: "HandleShutdownMsg",
			Handler:    _ServerMsgService_HandleShutdownMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpcMsg.proto",
}

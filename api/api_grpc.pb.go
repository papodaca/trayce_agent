// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/api.proto

package api

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

// TrayceAgentClient is the client API for TrayceAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrayceAgentClient interface {
	SendFlowsObserved(ctx context.Context, in *Flows, opts ...grpc.CallOption) (*Reply, error)
	SendAgentStarted(ctx context.Context, in *AgentStarted, opts ...grpc.CallOption) (*Reply, error)
	OpenCommandStream(ctx context.Context, opts ...grpc.CallOption) (TrayceAgent_OpenCommandStreamClient, error)
}

type trayceAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewTrayceAgentClient(cc grpc.ClientConnInterface) TrayceAgentClient {
	return &trayceAgentClient{cc}
}

func (c *trayceAgentClient) SendFlowsObserved(ctx context.Context, in *Flows, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/api.TrayceAgent/SendFlowsObserved", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trayceAgentClient) SendAgentStarted(ctx context.Context, in *AgentStarted, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/api.TrayceAgent/SendAgentStarted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trayceAgentClient) OpenCommandStream(ctx context.Context, opts ...grpc.CallOption) (TrayceAgent_OpenCommandStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrayceAgent_ServiceDesc.Streams[0], "/api.TrayceAgent/OpenCommandStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &trayceAgentOpenCommandStreamClient{stream}
	return x, nil
}

type TrayceAgent_OpenCommandStreamClient interface {
	Send(*NooP) error
	Recv() (*Command, error)
	grpc.ClientStream
}

type trayceAgentOpenCommandStreamClient struct {
	grpc.ClientStream
}

func (x *trayceAgentOpenCommandStreamClient) Send(m *NooP) error {
	return x.ClientStream.SendMsg(m)
}

func (x *trayceAgentOpenCommandStreamClient) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrayceAgentServer is the server API for TrayceAgent service.
// All implementations must embed UnimplementedTrayceAgentServer
// for forward compatibility
type TrayceAgentServer interface {
	SendFlowsObserved(context.Context, *Flows) (*Reply, error)
	SendAgentStarted(context.Context, *AgentStarted) (*Reply, error)
	OpenCommandStream(TrayceAgent_OpenCommandStreamServer) error
	mustEmbedUnimplementedTrayceAgentServer()
}

// UnimplementedTrayceAgentServer must be embedded to have forward compatible implementations.
type UnimplementedTrayceAgentServer struct {
}

func (UnimplementedTrayceAgentServer) SendFlowsObserved(context.Context, *Flows) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFlowsObserved not implemented")
}
func (UnimplementedTrayceAgentServer) SendAgentStarted(context.Context, *AgentStarted) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAgentStarted not implemented")
}
func (UnimplementedTrayceAgentServer) OpenCommandStream(TrayceAgent_OpenCommandStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenCommandStream not implemented")
}
func (UnimplementedTrayceAgentServer) mustEmbedUnimplementedTrayceAgentServer() {}

// UnsafeTrayceAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrayceAgentServer will
// result in compilation errors.
type UnsafeTrayceAgentServer interface {
	mustEmbedUnimplementedTrayceAgentServer()
}

func RegisterTrayceAgentServer(s grpc.ServiceRegistrar, srv TrayceAgentServer) {
	s.RegisterService(&TrayceAgent_ServiceDesc, srv)
}

func _TrayceAgent_SendFlowsObserved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Flows)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrayceAgentServer).SendFlowsObserved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TrayceAgent/SendFlowsObserved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrayceAgentServer).SendFlowsObserved(ctx, req.(*Flows))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrayceAgent_SendAgentStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentStarted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrayceAgentServer).SendAgentStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TrayceAgent/SendAgentStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrayceAgentServer).SendAgentStarted(ctx, req.(*AgentStarted))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrayceAgent_OpenCommandStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TrayceAgentServer).OpenCommandStream(&trayceAgentOpenCommandStreamServer{stream})
}

type TrayceAgent_OpenCommandStreamServer interface {
	Send(*Command) error
	Recv() (*NooP, error)
	grpc.ServerStream
}

type trayceAgentOpenCommandStreamServer struct {
	grpc.ServerStream
}

func (x *trayceAgentOpenCommandStreamServer) Send(m *Command) error {
	return x.ServerStream.SendMsg(m)
}

func (x *trayceAgentOpenCommandStreamServer) Recv() (*NooP, error) {
	m := new(NooP)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrayceAgent_ServiceDesc is the grpc.ServiceDesc for TrayceAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrayceAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TrayceAgent",
	HandlerType: (*TrayceAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendFlowsObserved",
			Handler:    _TrayceAgent_SendFlowsObserved_Handler,
		},
		{
			MethodName: "SendAgentStarted",
			Handler:    _TrayceAgent_SendAgentStarted_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenCommandStream",
			Handler:       _TrayceAgent_OpenCommandStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/api.proto",
}

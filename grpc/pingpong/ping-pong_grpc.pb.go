// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: pingpong/ping-pong.proto

package pingpong

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

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongClient interface {
	RpcPing(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type pingPongClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongClient(cc grpc.ClientConnInterface) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) RpcPing(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pingpong.PingPong/RpcPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServer is the server API for PingPong service.
// All implementations must embed UnimplementedPingPongServer
// for forward compatibility
type PingPongServer interface {
	RpcPing(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedPingPongServer()
}

// UnimplementedPingPongServer must be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (UnimplementedPingPongServer) RpcPing(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcPing not implemented")
}
func (UnimplementedPingPongServer) mustEmbedUnimplementedPingPongServer() {}

// UnsafePingPongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingPongServer will
// result in compilation errors.
type UnsafePingPongServer interface {
	mustEmbedUnimplementedPingPongServer()
}

func RegisterPingPongServer(s grpc.ServiceRegistrar, srv PingPongServer) {
	s.RegisterService(&PingPong_ServiceDesc, srv)
}

func _PingPong_RpcPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).RpcPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PingPong/RpcPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).RpcPing(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// PingPong_ServiceDesc is the grpc.ServiceDesc for PingPong service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingPong_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcPing",
			Handler:    _PingPong_RpcPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pingpong/ping-pong.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	SomeMethodUnderAuth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SomeRes, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) SomeMethodUnderAuth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SomeRes, error) {
	out := new(SomeRes)
	err := c.cc.Invoke(ctx, "/Api.Game/SomeMethodUnderAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations should embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	SomeMethodUnderAuth(context.Context, *emptypb.Empty) (*SomeRes, error)
}

// UnimplementedGameServer should be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) SomeMethodUnderAuth(context.Context, *emptypb.Empty) (*SomeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SomeMethodUnderAuth not implemented")
}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_SomeMethodUnderAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SomeMethodUnderAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Api.Game/SomeMethodUnderAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SomeMethodUnderAuth(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Api.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SomeMethodUnderAuth",
			Handler:    _Game_SomeMethodUnderAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/api/proto/game.proto",
}

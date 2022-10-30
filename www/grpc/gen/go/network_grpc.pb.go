// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: network.proto

package pactus

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

// NetworkClient is the client API for Network service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkClient interface {
	GetNetworkInfo(ctx context.Context, in *NetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error)
	GetPeerInfo(ctx context.Context, in *PeerInfoRequest, opts ...grpc.CallOption) (*PeerInfoResponse, error)
}

type networkClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkClient(cc grpc.ClientConnInterface) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) GetNetworkInfo(ctx context.Context, in *NetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error) {
	out := new(NetworkInfoResponse)
	err := c.cc.Invoke(ctx, "/pactus.Network/GetNetworkInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetPeerInfo(ctx context.Context, in *PeerInfoRequest, opts ...grpc.CallOption) (*PeerInfoResponse, error) {
	out := new(PeerInfoResponse)
	err := c.cc.Invoke(ctx, "/pactus.Network/GetPeerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServer is the server API for Network service.
// All implementations should embed UnimplementedNetworkServer
// for forward compatibility
type NetworkServer interface {
	GetNetworkInfo(context.Context, *NetworkInfoRequest) (*NetworkInfoResponse, error)
	GetPeerInfo(context.Context, *PeerInfoRequest) (*PeerInfoResponse, error)
}

// UnimplementedNetworkServer should be embedded to have forward compatible implementations.
type UnimplementedNetworkServer struct {
}

func (UnimplementedNetworkServer) GetNetworkInfo(context.Context, *NetworkInfoRequest) (*NetworkInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkInfo not implemented")
}
func (UnimplementedNetworkServer) GetPeerInfo(context.Context, *PeerInfoRequest) (*PeerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeerInfo not implemented")
}

// UnsafeNetworkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServer will
// result in compilation errors.
type UnsafeNetworkServer interface {
	mustEmbedUnimplementedNetworkServer()
}

func RegisterNetworkServer(s grpc.ServiceRegistrar, srv NetworkServer) {
	s.RegisterService(&Network_ServiceDesc, srv)
}

func _Network_GetNetworkInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetNetworkInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Network/GetNetworkInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetNetworkInfo(ctx, req.(*NetworkInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetPeerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetPeerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Network/GetPeerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetPeerInfo(ctx, req.(*PeerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Network_ServiceDesc is the grpc.ServiceDesc for Network service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Network_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pactus.Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkInfo",
			Handler:    _Network_GetNetworkInfo_Handler,
		},
		{
			MethodName: "GetPeerInfo",
			Handler:    _Network_GetPeerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network.proto",
}

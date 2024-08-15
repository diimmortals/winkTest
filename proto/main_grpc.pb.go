// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: proto/main.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Balancer_Redirect_FullMethodName = "/balancer.Balancer/Redirect"
)

// BalancerClient is the client API for Balancer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalancerClient interface {
	Redirect(ctx context.Context, in *RedirectRequest, opts ...grpc.CallOption) (*RedirectResponse, error)
}

type balancerClient struct {
	cc grpc.ClientConnInterface
}

func NewBalancerClient(cc grpc.ClientConnInterface) BalancerClient {
	return &balancerClient{cc}
}

func (c *balancerClient) Redirect(ctx context.Context, in *RedirectRequest, opts ...grpc.CallOption) (*RedirectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RedirectResponse)
	err := c.cc.Invoke(ctx, Balancer_Redirect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalancerServer is the server API for Balancer service.
// All implementations must embed UnimplementedBalancerServer
// for forward compatibility.
type BalancerServer interface {
	Redirect(context.Context, *RedirectRequest) (*RedirectResponse, error)
	mustEmbedUnimplementedBalancerServer()
}

// UnimplementedBalancerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBalancerServer struct{}

func (UnimplementedBalancerServer) Redirect(context.Context, *RedirectRequest) (*RedirectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redirect not implemented")
}
func (UnimplementedBalancerServer) mustEmbedUnimplementedBalancerServer() {}
func (UnimplementedBalancerServer) testEmbeddedByValue()                  {}

// UnsafeBalancerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalancerServer will
// result in compilation errors.
type UnsafeBalancerServer interface {
	mustEmbedUnimplementedBalancerServer()
}

func RegisterBalancerServer(s grpc.ServiceRegistrar, srv BalancerServer) {
	// If the following call pancis, it indicates UnimplementedBalancerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Balancer_ServiceDesc, srv)
}

func _Balancer_Redirect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).Redirect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Balancer_Redirect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).Redirect(ctx, req.(*RedirectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Balancer_ServiceDesc is the grpc.ServiceDesc for Balancer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Balancer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "balancer.Balancer",
	HandlerType: (*BalancerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Redirect",
			Handler:    _Balancer_Redirect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/main.proto",
}

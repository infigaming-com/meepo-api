// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: wallet/service/v1/wallet_event.proto

package v1

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
	WalletEvent_Event_FullMethodName = "/api.wallet.service.v1.WalletEvent/Event"
)

// WalletEventClient is the client API for WalletEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletEventClient interface {
	Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type walletEventClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletEventClient(cc grpc.ClientConnInterface) WalletEventClient {
	return &walletEventClient{cc}
}

func (c *walletEventClient) Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, WalletEvent_Event_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletEventServer is the server API for WalletEvent service.
// All implementations must embed UnimplementedWalletEventServer
// for forward compatibility.
type WalletEventServer interface {
	Event(context.Context, *EventRequest) (*EventResponse, error)
	mustEmbedUnimplementedWalletEventServer()
}

// UnimplementedWalletEventServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletEventServer struct{}

func (UnimplementedWalletEventServer) Event(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Event not implemented")
}
func (UnimplementedWalletEventServer) mustEmbedUnimplementedWalletEventServer() {}
func (UnimplementedWalletEventServer) testEmbeddedByValue()                     {}

// UnsafeWalletEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletEventServer will
// result in compilation errors.
type UnsafeWalletEventServer interface {
	mustEmbedUnimplementedWalletEventServer()
}

func RegisterWalletEventServer(s grpc.ServiceRegistrar, srv WalletEventServer) {
	// If the following call pancis, it indicates UnimplementedWalletEventServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WalletEvent_ServiceDesc, srv)
}

func _WalletEvent_Event_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletEventServer).Event(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletEvent_Event_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletEventServer).Event(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletEvent_ServiceDesc is the grpc.ServiceDesc for WalletEvent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletEvent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.wallet.service.v1.WalletEvent",
	HandlerType: (*WalletEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Event",
			Handler:    _WalletEvent_Event_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/service/v1/wallet_event.proto",
}

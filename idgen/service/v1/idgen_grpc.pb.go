// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: idgen/service/v1/idgen.proto

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
	Idgen_GenerateId_FullMethodName      = "/idgen.service.v1.Idgen/GenerateId"
	Idgen_BatchGenerateId_FullMethodName = "/idgen.service.v1.Idgen/BatchGenerateId"
)

// IdgenClient is the client API for Idgen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdgenClient interface {
	// 生成单个ID
	GenerateId(ctx context.Context, in *GenerateIdRequest, opts ...grpc.CallOption) (*GenerateIdReply, error)
	// 批量生成ID
	BatchGenerateId(ctx context.Context, in *BatchGenerateIdRequest, opts ...grpc.CallOption) (*BatchGenerateIdReply, error)
}

type idgenClient struct {
	cc grpc.ClientConnInterface
}

func NewIdgenClient(cc grpc.ClientConnInterface) IdgenClient {
	return &idgenClient{cc}
}

func (c *idgenClient) GenerateId(ctx context.Context, in *GenerateIdRequest, opts ...grpc.CallOption) (*GenerateIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateIdReply)
	err := c.cc.Invoke(ctx, Idgen_GenerateId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idgenClient) BatchGenerateId(ctx context.Context, in *BatchGenerateIdRequest, opts ...grpc.CallOption) (*BatchGenerateIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGenerateIdReply)
	err := c.cc.Invoke(ctx, Idgen_BatchGenerateId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdgenServer is the server API for Idgen service.
// All implementations must embed UnimplementedIdgenServer
// for forward compatibility.
type IdgenServer interface {
	// 生成单个ID
	GenerateId(context.Context, *GenerateIdRequest) (*GenerateIdReply, error)
	// 批量生成ID
	BatchGenerateId(context.Context, *BatchGenerateIdRequest) (*BatchGenerateIdReply, error)
	mustEmbedUnimplementedIdgenServer()
}

// UnimplementedIdgenServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIdgenServer struct{}

func (UnimplementedIdgenServer) GenerateId(context.Context, *GenerateIdRequest) (*GenerateIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateId not implemented")
}
func (UnimplementedIdgenServer) BatchGenerateId(context.Context, *BatchGenerateIdRequest) (*BatchGenerateIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGenerateId not implemented")
}
func (UnimplementedIdgenServer) mustEmbedUnimplementedIdgenServer() {}
func (UnimplementedIdgenServer) testEmbeddedByValue()               {}

// UnsafeIdgenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdgenServer will
// result in compilation errors.
type UnsafeIdgenServer interface {
	mustEmbedUnimplementedIdgenServer()
}

func RegisterIdgenServer(s grpc.ServiceRegistrar, srv IdgenServer) {
	// If the following call pancis, it indicates UnimplementedIdgenServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Idgen_ServiceDesc, srv)
}

func _Idgen_GenerateId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdgenServer).GenerateId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Idgen_GenerateId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdgenServer).GenerateId(ctx, req.(*GenerateIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idgen_BatchGenerateId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGenerateIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdgenServer).BatchGenerateId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Idgen_BatchGenerateId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdgenServer).BatchGenerateId(ctx, req.(*BatchGenerateIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Idgen_ServiceDesc is the grpc.ServiceDesc for Idgen service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Idgen_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idgen.service.v1.Idgen",
	HandlerType: (*IdgenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateId",
			Handler:    _Idgen_GenerateId_Handler,
		},
		{
			MethodName: "BatchGenerateId",
			Handler:    _Idgen_BatchGenerateId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idgen/service/v1/idgen.proto",
}

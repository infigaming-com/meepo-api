// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: review/service/v1/review.proto

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
	Review_CreateWithdraw_FullMethodName         = "/api.review.service.v1.Review/CreateWithdraw"
	Review_CreateOperatorWithdraw_FullMethodName = "/api.review.service.v1.Review/CreateOperatorWithdraw"
	Review_ReviewTicket_FullMethodName           = "/api.review.service.v1.Review/ReviewTicket"
	Review_AddComment_FullMethodName             = "/api.review.service.v1.Review/AddComment"
	Review_CancelTicket_FullMethodName           = "/api.review.service.v1.Review/CancelTicket"
	Review_ListTickets_FullMethodName            = "/api.review.service.v1.Review/ListTickets"
	Review_ListOperatorTickets_FullMethodName    = "/api.review.service.v1.Review/ListOperatorTickets"
	Review_GetTicket_FullMethodName              = "/api.review.service.v1.Review/GetTicket"
	Review_GetOperatorTicket_FullMethodName      = "/api.review.service.v1.Review/GetOperatorTicket"
)

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Review service provides review functionality.
type ReviewClient interface {
	CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawResponse, error)
	CreateOperatorWithdraw(ctx context.Context, in *CreateOperatorWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawResponse, error)
	ReviewTicket(ctx context.Context, in *ReviewTicketRequest, opts ...grpc.CallOption) (*ReviewTicketResponse, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error)
	// CancelTicket is used to manually cancel a ticket.
	// Ticket is cancellable only in paying status.
	CancelTicket(ctx context.Context, in *CancelTicketRequest, opts ...grpc.CallOption) (*CancelTicketResponse, error)
	ListTickets(ctx context.Context, in *ListTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error)
	ListOperatorTickets(ctx context.Context, in *ListOperatorTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error)
	GetOperatorTicket(ctx context.Context, in *GetOperatorTicketRequest, opts ...grpc.CallOption) (*GetOperatorTicketResponse, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWithdrawResponse)
	err := c.cc.Invoke(ctx, Review_CreateWithdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) CreateOperatorWithdraw(ctx context.Context, in *CreateOperatorWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWithdrawResponse)
	err := c.cc.Invoke(ctx, Review_CreateOperatorWithdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ReviewTicket(ctx context.Context, in *ReviewTicketRequest, opts ...grpc.CallOption) (*ReviewTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReviewTicketResponse)
	err := c.cc.Invoke(ctx, Review_ReviewTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCommentResponse)
	err := c.cc.Invoke(ctx, Review_AddComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) CancelTicket(ctx context.Context, in *CancelTicketRequest, opts ...grpc.CallOption) (*CancelTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelTicketResponse)
	err := c.cc.Invoke(ctx, Review_CancelTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ListTickets(ctx context.Context, in *ListTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTicketsResponse)
	err := c.cc.Invoke(ctx, Review_ListTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ListOperatorTickets(ctx context.Context, in *ListOperatorTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTicketsResponse)
	err := c.cc.Invoke(ctx, Review_ListOperatorTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTicketResponse)
	err := c.cc.Invoke(ctx, Review_GetTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) GetOperatorTicket(ctx context.Context, in *GetOperatorTicketRequest, opts ...grpc.CallOption) (*GetOperatorTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOperatorTicketResponse)
	err := c.cc.Invoke(ctx, Review_GetOperatorTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
// All implementations must embed UnimplementedReviewServer
// for forward compatibility.
//
// Review service provides review functionality.
type ReviewServer interface {
	CreateWithdraw(context.Context, *CreateWithdrawRequest) (*CreateWithdrawResponse, error)
	CreateOperatorWithdraw(context.Context, *CreateOperatorWithdrawRequest) (*CreateWithdrawResponse, error)
	ReviewTicket(context.Context, *ReviewTicketRequest) (*ReviewTicketResponse, error)
	AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error)
	// CancelTicket is used to manually cancel a ticket.
	// Ticket is cancellable only in paying status.
	CancelTicket(context.Context, *CancelTicketRequest) (*CancelTicketResponse, error)
	ListTickets(context.Context, *ListTicketsRequest) (*ListTicketsResponse, error)
	ListOperatorTickets(context.Context, *ListOperatorTicketsRequest) (*ListTicketsResponse, error)
	GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error)
	GetOperatorTicket(context.Context, *GetOperatorTicketRequest) (*GetOperatorTicketResponse, error)
	mustEmbedUnimplementedReviewServer()
}

// UnimplementedReviewServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReviewServer struct{}

func (UnimplementedReviewServer) CreateWithdraw(context.Context, *CreateWithdrawRequest) (*CreateWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWithdraw not implemented")
}
func (UnimplementedReviewServer) CreateOperatorWithdraw(context.Context, *CreateOperatorWithdrawRequest) (*CreateWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOperatorWithdraw not implemented")
}
func (UnimplementedReviewServer) ReviewTicket(context.Context, *ReviewTicketRequest) (*ReviewTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewTicket not implemented")
}
func (UnimplementedReviewServer) AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedReviewServer) CancelTicket(context.Context, *CancelTicketRequest) (*CancelTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTicket not implemented")
}
func (UnimplementedReviewServer) ListTickets(context.Context, *ListTicketsRequest) (*ListTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTickets not implemented")
}
func (UnimplementedReviewServer) ListOperatorTickets(context.Context, *ListOperatorTicketsRequest) (*ListTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperatorTickets not implemented")
}
func (UnimplementedReviewServer) GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedReviewServer) GetOperatorTicket(context.Context, *GetOperatorTicketRequest) (*GetOperatorTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorTicket not implemented")
}
func (UnimplementedReviewServer) mustEmbedUnimplementedReviewServer() {}
func (UnimplementedReviewServer) testEmbeddedByValue()                {}

// UnsafeReviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServer will
// result in compilation errors.
type UnsafeReviewServer interface {
	mustEmbedUnimplementedReviewServer()
}

func RegisterReviewServer(s grpc.ServiceRegistrar, srv ReviewServer) {
	// If the following call pancis, it indicates UnimplementedReviewServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Review_ServiceDesc, srv)
}

func _Review_CreateWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).CreateWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_CreateWithdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).CreateWithdraw(ctx, req.(*CreateWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_CreateOperatorWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOperatorWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).CreateOperatorWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_CreateOperatorWithdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).CreateOperatorWithdraw(ctx, req.(*CreateOperatorWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ReviewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ReviewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ReviewTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ReviewTicket(ctx, req.(*ReviewTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_CancelTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).CancelTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_CancelTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).CancelTicket(ctx, req.(*CancelTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ListTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ListTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ListTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ListTickets(ctx, req.(*ListTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ListOperatorTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperatorTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ListOperatorTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ListOperatorTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ListOperatorTickets(ctx, req.(*ListOperatorTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_GetTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_GetOperatorTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatorTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).GetOperatorTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_GetOperatorTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).GetOperatorTicket(ctx, req.(*GetOperatorTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Review_ServiceDesc is the grpc.ServiceDesc for Review service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Review_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.review.service.v1.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWithdraw",
			Handler:    _Review_CreateWithdraw_Handler,
		},
		{
			MethodName: "CreateOperatorWithdraw",
			Handler:    _Review_CreateOperatorWithdraw_Handler,
		},
		{
			MethodName: "ReviewTicket",
			Handler:    _Review_ReviewTicket_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _Review_AddComment_Handler,
		},
		{
			MethodName: "CancelTicket",
			Handler:    _Review_CancelTicket_Handler,
		},
		{
			MethodName: "ListTickets",
			Handler:    _Review_ListTickets_Handler,
		},
		{
			MethodName: "ListOperatorTickets",
			Handler:    _Review_ListOperatorTickets_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _Review_GetTicket_Handler,
		},
		{
			MethodName: "GetOperatorTicket",
			Handler:    _Review_GetOperatorTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review/service/v1/review.proto",
}

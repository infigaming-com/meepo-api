// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: backoffice/service/v1/backoffice_report.proto

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
	BackofficeReport_GetSummary_FullMethodName               = "/api.backoffice.service.v1.BackofficeReport/GetSummary"
	BackofficeReport_ListSummaries_FullMethodName            = "/api.backoffice.service.v1.BackofficeReport/ListSummaries"
	BackofficeReport_GetGameDataSummary_FullMethodName       = "/api.backoffice.service.v1.BackofficeReport/GetGameDataSummary"
	BackofficeReport_ListGameData_FullMethodName             = "/api.backoffice.service.v1.BackofficeReport/ListGameData"
	BackofficeReport_GetPlayerGameDataSummary_FullMethodName = "/api.backoffice.service.v1.BackofficeReport/GetPlayerGameDataSummary"
	BackofficeReport_ListPlayerGameData_FullMethodName       = "/api.backoffice.service.v1.BackofficeReport/ListPlayerGameData"
	BackofficeReport_GetDepositSummaries_FullMethodName      = "/api.backoffice.service.v1.BackofficeReport/GetDepositSummaries"
	BackofficeReport_ListDepositDetails_FullMethodName       = "/api.backoffice.service.v1.BackofficeReport/ListDepositDetails"
	BackofficeReport_GetWithdrawSummaries_FullMethodName     = "/api.backoffice.service.v1.BackofficeReport/GetWithdrawSummaries"
	BackofficeReport_ListWithdrawDetails_FullMethodName      = "/api.backoffice.service.v1.BackofficeReport/ListWithdrawDetails"
	BackofficeReport_ListRegisterRetention_FullMethodName    = "/api.backoffice.service.v1.BackofficeReport/ListRegisterRetention"
	BackofficeReport_ListDepositVtgDetails_FullMethodName    = "/api.backoffice.service.v1.BackofficeReport/ListDepositVtgDetails"
	BackofficeReport_ListWithdrawVtgDetails_FullMethodName   = "/api.backoffice.service.v1.BackofficeReport/ListWithdrawVtgDetails"
)

// BackofficeReportClient is the client API for BackofficeReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackofficeReportClient interface {
	GetSummary(ctx context.Context, in *GetSummaryRequest, opts ...grpc.CallOption) (*GetSummaryResponse, error)
	ListSummaries(ctx context.Context, in *ListSummariesRequest, opts ...grpc.CallOption) (*ListSummariesResponse, error)
	GetGameDataSummary(ctx context.Context, in *GetGameSummaryRequest, opts ...grpc.CallOption) (*GetGameSummaryResponse, error)
	ListGameData(ctx context.Context, in *GetGameDataRequest, opts ...grpc.CallOption) (*GetGameDataResponse, error)
	GetPlayerGameDataSummary(ctx context.Context, in *GetPlayerGameSummaryRequest, opts ...grpc.CallOption) (*GetPlayerGameSummaryResponse, error)
	ListPlayerGameData(ctx context.Context, in *GetPlayerGameDataRequest, opts ...grpc.CallOption) (*GetPlayerGameDataResponse, error)
	GetDepositSummaries(ctx context.Context, in *GetDepositSummariesRequest, opts ...grpc.CallOption) (*GetDepositSummariesResponse, error)
	ListDepositDetails(ctx context.Context, in *ListDepositDetailsRequest, opts ...grpc.CallOption) (*ListDepositDetailsResponse, error)
	GetWithdrawSummaries(ctx context.Context, in *GetWithdrawSummariesRequest, opts ...grpc.CallOption) (*GetWithdrawSummariesResponse, error)
	ListWithdrawDetails(ctx context.Context, in *ListWithdrawDetailsRequest, opts ...grpc.CallOption) (*ListWithdrawDetailsResponse, error)
	ListRegisterRetention(ctx context.Context, in *ListRegisterRetentionRequest, opts ...grpc.CallOption) (*ListRegisterRetentionResponse, error)
	ListDepositVtgDetails(ctx context.Context, in *ListDepositVtgDetailsRequest, opts ...grpc.CallOption) (*ListDepositVtgDetailsResponse, error)
	ListWithdrawVtgDetails(ctx context.Context, in *ListWithdrawVtgDetailsRequest, opts ...grpc.CallOption) (*ListWithdrawVtgDetailsResponse, error)
}

type backofficeReportClient struct {
	cc grpc.ClientConnInterface
}

func NewBackofficeReportClient(cc grpc.ClientConnInterface) BackofficeReportClient {
	return &backofficeReportClient{cc}
}

func (c *backofficeReportClient) GetSummary(ctx context.Context, in *GetSummaryRequest, opts ...grpc.CallOption) (*GetSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSummaryResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_GetSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListSummaries(ctx context.Context, in *ListSummariesRequest, opts ...grpc.CallOption) (*ListSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSummariesResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) GetGameDataSummary(ctx context.Context, in *GetGameSummaryRequest, opts ...grpc.CallOption) (*GetGameSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGameSummaryResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_GetGameDataSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListGameData(ctx context.Context, in *GetGameDataRequest, opts ...grpc.CallOption) (*GetGameDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGameDataResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListGameData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) GetPlayerGameDataSummary(ctx context.Context, in *GetPlayerGameSummaryRequest, opts ...grpc.CallOption) (*GetPlayerGameSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlayerGameSummaryResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_GetPlayerGameDataSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListPlayerGameData(ctx context.Context, in *GetPlayerGameDataRequest, opts ...grpc.CallOption) (*GetPlayerGameDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlayerGameDataResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListPlayerGameData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) GetDepositSummaries(ctx context.Context, in *GetDepositSummariesRequest, opts ...grpc.CallOption) (*GetDepositSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDepositSummariesResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_GetDepositSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListDepositDetails(ctx context.Context, in *ListDepositDetailsRequest, opts ...grpc.CallOption) (*ListDepositDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDepositDetailsResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListDepositDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) GetWithdrawSummaries(ctx context.Context, in *GetWithdrawSummariesRequest, opts ...grpc.CallOption) (*GetWithdrawSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWithdrawSummariesResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_GetWithdrawSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListWithdrawDetails(ctx context.Context, in *ListWithdrawDetailsRequest, opts ...grpc.CallOption) (*ListWithdrawDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListWithdrawDetailsResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListWithdrawDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListRegisterRetention(ctx context.Context, in *ListRegisterRetentionRequest, opts ...grpc.CallOption) (*ListRegisterRetentionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRegisterRetentionResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListRegisterRetention_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListDepositVtgDetails(ctx context.Context, in *ListDepositVtgDetailsRequest, opts ...grpc.CallOption) (*ListDepositVtgDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDepositVtgDetailsResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListDepositVtgDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backofficeReportClient) ListWithdrawVtgDetails(ctx context.Context, in *ListWithdrawVtgDetailsRequest, opts ...grpc.CallOption) (*ListWithdrawVtgDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListWithdrawVtgDetailsResponse)
	err := c.cc.Invoke(ctx, BackofficeReport_ListWithdrawVtgDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackofficeReportServer is the server API for BackofficeReport service.
// All implementations must embed UnimplementedBackofficeReportServer
// for forward compatibility.
type BackofficeReportServer interface {
	GetSummary(context.Context, *GetSummaryRequest) (*GetSummaryResponse, error)
	ListSummaries(context.Context, *ListSummariesRequest) (*ListSummariesResponse, error)
	GetGameDataSummary(context.Context, *GetGameSummaryRequest) (*GetGameSummaryResponse, error)
	ListGameData(context.Context, *GetGameDataRequest) (*GetGameDataResponse, error)
	GetPlayerGameDataSummary(context.Context, *GetPlayerGameSummaryRequest) (*GetPlayerGameSummaryResponse, error)
	ListPlayerGameData(context.Context, *GetPlayerGameDataRequest) (*GetPlayerGameDataResponse, error)
	GetDepositSummaries(context.Context, *GetDepositSummariesRequest) (*GetDepositSummariesResponse, error)
	ListDepositDetails(context.Context, *ListDepositDetailsRequest) (*ListDepositDetailsResponse, error)
	GetWithdrawSummaries(context.Context, *GetWithdrawSummariesRequest) (*GetWithdrawSummariesResponse, error)
	ListWithdrawDetails(context.Context, *ListWithdrawDetailsRequest) (*ListWithdrawDetailsResponse, error)
	ListRegisterRetention(context.Context, *ListRegisterRetentionRequest) (*ListRegisterRetentionResponse, error)
	ListDepositVtgDetails(context.Context, *ListDepositVtgDetailsRequest) (*ListDepositVtgDetailsResponse, error)
	ListWithdrawVtgDetails(context.Context, *ListWithdrawVtgDetailsRequest) (*ListWithdrawVtgDetailsResponse, error)
	mustEmbedUnimplementedBackofficeReportServer()
}

// UnimplementedBackofficeReportServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBackofficeReportServer struct{}

func (UnimplementedBackofficeReportServer) GetSummary(context.Context, *GetSummaryRequest) (*GetSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummary not implemented")
}
func (UnimplementedBackofficeReportServer) ListSummaries(context.Context, *ListSummariesRequest) (*ListSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSummaries not implemented")
}
func (UnimplementedBackofficeReportServer) GetGameDataSummary(context.Context, *GetGameSummaryRequest) (*GetGameSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameDataSummary not implemented")
}
func (UnimplementedBackofficeReportServer) ListGameData(context.Context, *GetGameDataRequest) (*GetGameDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGameData not implemented")
}
func (UnimplementedBackofficeReportServer) GetPlayerGameDataSummary(context.Context, *GetPlayerGameSummaryRequest) (*GetPlayerGameSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerGameDataSummary not implemented")
}
func (UnimplementedBackofficeReportServer) ListPlayerGameData(context.Context, *GetPlayerGameDataRequest) (*GetPlayerGameDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlayerGameData not implemented")
}
func (UnimplementedBackofficeReportServer) GetDepositSummaries(context.Context, *GetDepositSummariesRequest) (*GetDepositSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepositSummaries not implemented")
}
func (UnimplementedBackofficeReportServer) ListDepositDetails(context.Context, *ListDepositDetailsRequest) (*ListDepositDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepositDetails not implemented")
}
func (UnimplementedBackofficeReportServer) GetWithdrawSummaries(context.Context, *GetWithdrawSummariesRequest) (*GetWithdrawSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawSummaries not implemented")
}
func (UnimplementedBackofficeReportServer) ListWithdrawDetails(context.Context, *ListWithdrawDetailsRequest) (*ListWithdrawDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWithdrawDetails not implemented")
}
func (UnimplementedBackofficeReportServer) ListRegisterRetention(context.Context, *ListRegisterRetentionRequest) (*ListRegisterRetentionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegisterRetention not implemented")
}
func (UnimplementedBackofficeReportServer) ListDepositVtgDetails(context.Context, *ListDepositVtgDetailsRequest) (*ListDepositVtgDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepositVtgDetails not implemented")
}
func (UnimplementedBackofficeReportServer) ListWithdrawVtgDetails(context.Context, *ListWithdrawVtgDetailsRequest) (*ListWithdrawVtgDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWithdrawVtgDetails not implemented")
}
func (UnimplementedBackofficeReportServer) mustEmbedUnimplementedBackofficeReportServer() {}
func (UnimplementedBackofficeReportServer) testEmbeddedByValue()                          {}

// UnsafeBackofficeReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackofficeReportServer will
// result in compilation errors.
type UnsafeBackofficeReportServer interface {
	mustEmbedUnimplementedBackofficeReportServer()
}

func RegisterBackofficeReportServer(s grpc.ServiceRegistrar, srv BackofficeReportServer) {
	// If the following call pancis, it indicates UnimplementedBackofficeReportServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BackofficeReport_ServiceDesc, srv)
}

func _BackofficeReport_GetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).GetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_GetSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).GetSummary(ctx, req.(*GetSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListSummaries(ctx, req.(*ListSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_GetGameDataSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).GetGameDataSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_GetGameDataSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).GetGameDataSummary(ctx, req.(*GetGameSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListGameData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListGameData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListGameData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListGameData(ctx, req.(*GetGameDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_GetPlayerGameDataSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerGameSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).GetPlayerGameDataSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_GetPlayerGameDataSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).GetPlayerGameDataSummary(ctx, req.(*GetPlayerGameSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListPlayerGameData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerGameDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListPlayerGameData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListPlayerGameData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListPlayerGameData(ctx, req.(*GetPlayerGameDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_GetDepositSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepositSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).GetDepositSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_GetDepositSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).GetDepositSummaries(ctx, req.(*GetDepositSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListDepositDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepositDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListDepositDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListDepositDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListDepositDetails(ctx, req.(*ListDepositDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_GetWithdrawSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).GetWithdrawSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_GetWithdrawSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).GetWithdrawSummaries(ctx, req.(*GetWithdrawSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListWithdrawDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWithdrawDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListWithdrawDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListWithdrawDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListWithdrawDetails(ctx, req.(*ListWithdrawDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListRegisterRetention_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegisterRetentionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListRegisterRetention(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListRegisterRetention_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListRegisterRetention(ctx, req.(*ListRegisterRetentionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListDepositVtgDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepositVtgDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListDepositVtgDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListDepositVtgDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListDepositVtgDetails(ctx, req.(*ListDepositVtgDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackofficeReport_ListWithdrawVtgDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWithdrawVtgDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackofficeReportServer).ListWithdrawVtgDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackofficeReport_ListWithdrawVtgDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackofficeReportServer).ListWithdrawVtgDetails(ctx, req.(*ListWithdrawVtgDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BackofficeReport_ServiceDesc is the grpc.ServiceDesc for BackofficeReport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackofficeReport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.backoffice.service.v1.BackofficeReport",
	HandlerType: (*BackofficeReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSummary",
			Handler:    _BackofficeReport_GetSummary_Handler,
		},
		{
			MethodName: "ListSummaries",
			Handler:    _BackofficeReport_ListSummaries_Handler,
		},
		{
			MethodName: "GetGameDataSummary",
			Handler:    _BackofficeReport_GetGameDataSummary_Handler,
		},
		{
			MethodName: "ListGameData",
			Handler:    _BackofficeReport_ListGameData_Handler,
		},
		{
			MethodName: "GetPlayerGameDataSummary",
			Handler:    _BackofficeReport_GetPlayerGameDataSummary_Handler,
		},
		{
			MethodName: "ListPlayerGameData",
			Handler:    _BackofficeReport_ListPlayerGameData_Handler,
		},
		{
			MethodName: "GetDepositSummaries",
			Handler:    _BackofficeReport_GetDepositSummaries_Handler,
		},
		{
			MethodName: "ListDepositDetails",
			Handler:    _BackofficeReport_ListDepositDetails_Handler,
		},
		{
			MethodName: "GetWithdrawSummaries",
			Handler:    _BackofficeReport_GetWithdrawSummaries_Handler,
		},
		{
			MethodName: "ListWithdrawDetails",
			Handler:    _BackofficeReport_ListWithdrawDetails_Handler,
		},
		{
			MethodName: "ListRegisterRetention",
			Handler:    _BackofficeReport_ListRegisterRetention_Handler,
		},
		{
			MethodName: "ListDepositVtgDetails",
			Handler:    _BackofficeReport_ListDepositVtgDetails_Handler,
		},
		{
			MethodName: "ListWithdrawVtgDetails",
			Handler:    _BackofficeReport_ListWithdrawVtgDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backoffice/service/v1/backoffice_report.proto",
}

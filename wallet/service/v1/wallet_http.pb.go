// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v5.29.3
// source: wallet/service/v1/wallet.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationWalletGetUserBalance = "/api.wallet.service.v1.Wallet/GetUserBalance"
const OperationWalletHealthCheck = "/api.wallet.service.v1.Wallet/HealthCheck"

type WalletHTTPServer interface {
	GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error)
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterWalletHTTPServer(s *http.Server, srv WalletHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/wallet/healthcheck", _Wallet_HealthCheck1_HTTP_Handler(srv))
	r.GET("/v1/wallet/balance/{user_id}", _Wallet_GetUserBalance0_HTTP_Handler(srv))
}

func _Wallet_HealthCheck1_HTTP_Handler(srv WalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HealthCheckRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWalletHealthCheck)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HealthCheck(ctx, req.(*HealthCheckRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HealthCheckResponse)
		return ctx.Result(200, reply)
	}
}

func _Wallet_GetUserBalance0_HTTP_Handler(srv WalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserBalanceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWalletGetUserBalance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserBalance(ctx, req.(*GetUserBalanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserBalanceResponse)
		return ctx.Result(200, reply)
	}
}

type WalletHTTPClient interface {
	GetUserBalance(ctx context.Context, req *GetUserBalanceRequest, opts ...http.CallOption) (rsp *GetUserBalanceResponse, err error)
	HealthCheck(ctx context.Context, req *HealthCheckRequest, opts ...http.CallOption) (rsp *HealthCheckResponse, err error)
}

type WalletHTTPClientImpl struct {
	cc *http.Client
}

func NewWalletHTTPClient(client *http.Client) WalletHTTPClient {
	return &WalletHTTPClientImpl{client}
}

func (c *WalletHTTPClientImpl) GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...http.CallOption) (*GetUserBalanceResponse, error) {
	var out GetUserBalanceResponse
	pattern := "/v1/wallet/balance/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWalletGetUserBalance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WalletHTTPClientImpl) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...http.CallOption) (*HealthCheckResponse, error) {
	var out HealthCheckResponse
	pattern := "/v1/wallet/healthcheck"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWalletHealthCheck))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

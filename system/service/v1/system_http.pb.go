// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.25.6
// source: system/service/v1/system.proto

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

const OperationSystemAddCurrency = "/system.service.v1.System/AddCurrency"
const OperationSystemGetCurrencies = "/system.service.v1.System/GetCurrencies"
const OperationSystemHealthCheck = "/system.service.v1.System/HealthCheck"

type SystemHTTPServer interface {
	AddCurrency(context.Context, *AddCurrencyRequest) (*AddCurrencyResponse, error)
	GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterSystemHTTPServer(s *http.Server, srv SystemHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/system/healthcheck", _System_HealthCheck1_HTTP_Handler(srv))
	r.POST("/v1/system/add-currency", _System_AddCurrency0_HTTP_Handler(srv))
	r.POST("/v1/system/get-currencies", _System_GetCurrencies0_HTTP_Handler(srv))
}

func _System_HealthCheck1_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HealthCheckRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemHealthCheck)
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

func _System_AddCurrency0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCurrencyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAddCurrency)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddCurrency(ctx, req.(*AddCurrencyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddCurrencyResponse)
		return ctx.Result(200, reply)
	}
}

func _System_GetCurrencies0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrenciesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetCurrencies)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrencies(ctx, req.(*GetCurrenciesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrenciesResponse)
		return ctx.Result(200, reply)
	}
}

type SystemHTTPClient interface {
	AddCurrency(ctx context.Context, req *AddCurrencyRequest, opts ...http.CallOption) (rsp *AddCurrencyResponse, err error)
	GetCurrencies(ctx context.Context, req *GetCurrenciesRequest, opts ...http.CallOption) (rsp *GetCurrenciesResponse, err error)
	HealthCheck(ctx context.Context, req *HealthCheckRequest, opts ...http.CallOption) (rsp *HealthCheckResponse, err error)
}

type SystemHTTPClientImpl struct {
	cc *http.Client
}

func NewSystemHTTPClient(client *http.Client) SystemHTTPClient {
	return &SystemHTTPClientImpl{client}
}

func (c *SystemHTTPClientImpl) AddCurrency(ctx context.Context, in *AddCurrencyRequest, opts ...http.CallOption) (*AddCurrencyResponse, error) {
	var out AddCurrencyResponse
	pattern := "/v1/system/add-currency"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAddCurrency))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SystemHTTPClientImpl) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...http.CallOption) (*GetCurrenciesResponse, error) {
	var out GetCurrenciesResponse
	pattern := "/v1/system/get-currencies"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemGetCurrencies))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SystemHTTPClientImpl) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...http.CallOption) (*HealthCheckResponse, error) {
	var out HealthCheckResponse
	pattern := "/v1/system/healthcheck"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemHealthCheck))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

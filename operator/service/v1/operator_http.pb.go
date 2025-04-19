// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.25.6
// source: operator/service/v1/operator.proto

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

const OperationOperatorAddOperator = "/api.operator.service.v1.Operator/AddOperator"
const OperationOperatorGetOperatorCurrencies = "/api.operator.service.v1.Operator/GetOperatorCurrencies"
const OperationOperatorUpdateOperator = "/api.operator.service.v1.Operator/UpdateOperator"
const OperationOperatorUpdateOperatorCurrency = "/api.operator.service.v1.Operator/UpdateOperatorCurrency"

type OperatorHTTPServer interface {
	AddOperator(context.Context, *AddOperatorRequest) (*AddOperatorResponse, error)
	GetOperatorCurrencies(context.Context, *GetOperatorCurrenciesRequest) (*GetOperatorCurrenciesResponse, error)
	UpdateOperator(context.Context, *UpdateOperatorRequest) (*UpdateOperatorResponse, error)
	UpdateOperatorCurrency(context.Context, *UpdateOperatorCurrencyRequest) (*UpdateOperatorCurrencyResponse, error)
}

func RegisterOperatorHTTPServer(s *http.Server, srv OperatorHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/operator/add", _Operator_AddOperator0_HTTP_Handler(srv))
	r.POST("/v1/operator/update", _Operator_UpdateOperator0_HTTP_Handler(srv))
	r.POST("/v1/operator/currencies/update", _Operator_UpdateOperatorCurrency0_HTTP_Handler(srv))
	r.POST("/v1/operator/currencies/get", _Operator_GetOperatorCurrencies0_HTTP_Handler(srv))
}

func _Operator_AddOperator0_HTTP_Handler(srv OperatorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddOperatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperatorAddOperator)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddOperator(ctx, req.(*AddOperatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddOperatorResponse)
		return ctx.Result(200, reply)
	}
}

func _Operator_UpdateOperator0_HTTP_Handler(srv OperatorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOperatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperatorUpdateOperator)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOperator(ctx, req.(*UpdateOperatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateOperatorResponse)
		return ctx.Result(200, reply)
	}
}

func _Operator_UpdateOperatorCurrency0_HTTP_Handler(srv OperatorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOperatorCurrencyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperatorUpdateOperatorCurrency)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOperatorCurrency(ctx, req.(*UpdateOperatorCurrencyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateOperatorCurrencyResponse)
		return ctx.Result(200, reply)
	}
}

func _Operator_GetOperatorCurrencies0_HTTP_Handler(srv OperatorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOperatorCurrenciesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperatorGetOperatorCurrencies)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOperatorCurrencies(ctx, req.(*GetOperatorCurrenciesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOperatorCurrenciesResponse)
		return ctx.Result(200, reply)
	}
}

type OperatorHTTPClient interface {
	AddOperator(ctx context.Context, req *AddOperatorRequest, opts ...http.CallOption) (rsp *AddOperatorResponse, err error)
	GetOperatorCurrencies(ctx context.Context, req *GetOperatorCurrenciesRequest, opts ...http.CallOption) (rsp *GetOperatorCurrenciesResponse, err error)
	UpdateOperator(ctx context.Context, req *UpdateOperatorRequest, opts ...http.CallOption) (rsp *UpdateOperatorResponse, err error)
	UpdateOperatorCurrency(ctx context.Context, req *UpdateOperatorCurrencyRequest, opts ...http.CallOption) (rsp *UpdateOperatorCurrencyResponse, err error)
}

type OperatorHTTPClientImpl struct {
	cc *http.Client
}

func NewOperatorHTTPClient(client *http.Client) OperatorHTTPClient {
	return &OperatorHTTPClientImpl{client}
}

func (c *OperatorHTTPClientImpl) AddOperator(ctx context.Context, in *AddOperatorRequest, opts ...http.CallOption) (*AddOperatorResponse, error) {
	var out AddOperatorResponse
	pattern := "/v1/operator/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperatorAddOperator))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperatorHTTPClientImpl) GetOperatorCurrencies(ctx context.Context, in *GetOperatorCurrenciesRequest, opts ...http.CallOption) (*GetOperatorCurrenciesResponse, error) {
	var out GetOperatorCurrenciesResponse
	pattern := "/v1/operator/currencies/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperatorGetOperatorCurrencies))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperatorHTTPClientImpl) UpdateOperator(ctx context.Context, in *UpdateOperatorRequest, opts ...http.CallOption) (*UpdateOperatorResponse, error) {
	var out UpdateOperatorResponse
	pattern := "/v1/operator/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperatorUpdateOperator))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperatorHTTPClientImpl) UpdateOperatorCurrency(ctx context.Context, in *UpdateOperatorCurrencyRequest, opts ...http.CallOption) (*UpdateOperatorCurrencyResponse, error) {
	var out UpdateOperatorCurrencyResponse
	pattern := "/v1/operator/currencies/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperatorUpdateOperatorCurrency))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

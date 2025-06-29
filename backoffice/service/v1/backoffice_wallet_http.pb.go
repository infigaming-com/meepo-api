// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.25.6
// source: backoffice/service/v1/backoffice_wallet.proto

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

const OperationBackofficeWalletAddWalletCurrency = "/api.backoffice.service.v1.BackofficeWallet/AddWalletCurrency"
const OperationBackofficeWalletGetWalletCreditTransactions = "/api.backoffice.service.v1.BackofficeWallet/GetWalletCreditTransactions"
const OperationBackofficeWalletGetWalletCredits = "/api.backoffice.service.v1.BackofficeWallet/GetWalletCredits"
const OperationBackofficeWalletGetWallets = "/api.backoffice.service.v1.BackofficeWallet/GetWallets"
const OperationBackofficeWalletListWalletBalanceTransactions = "/api.backoffice.service.v1.BackofficeWallet/ListWalletBalanceTransactions"
const OperationBackofficeWalletListWalletCurrencies = "/api.backoffice.service.v1.BackofficeWallet/ListWalletCurrencies"
const OperationBackofficeWalletUpdateWallet = "/api.backoffice.service.v1.BackofficeWallet/UpdateWallet"
const OperationBackofficeWalletUpdateWalletCurrency = "/api.backoffice.service.v1.BackofficeWallet/UpdateWalletCurrency"

type BackofficeWalletHTTPServer interface {
	AddWalletCurrency(context.Context, *AddWalletCurrencyRequest) (*AddWalletCurrencyResponse, error)
	GetWalletCreditTransactions(context.Context, *GetWalletCreditTransactionsRequest) (*GetWalletCreditTransactionsResponse, error)
	GetWalletCredits(context.Context, *GetWalletCreditsRequest) (*GetWalletCreditsResponse, error)
	GetWallets(context.Context, *GetWalletsRequest) (*GetWalletsResponse, error)
	// ListWalletBalanceTransactions ListWalletBalanceTransactions provides balance transactions for a specific user in User transactions page.
	ListWalletBalanceTransactions(context.Context, *ListWalletBalanceTransactionsRequest) (*ListWalletBalanceTransactionsResponse, error)
	ListWalletCurrencies(context.Context, *ListWalletCurrenciesRequest) (*ListWalletCurrenciesResponse, error)
	UpdateWallet(context.Context, *UpdateWalletRequest) (*UpdateWalletResponse, error)
	UpdateWalletCurrency(context.Context, *UpdateWalletCurrencyRequest) (*UpdateWalletCurrencyResponse, error)
}

func RegisterBackofficeWalletHTTPServer(s *http.Server, srv BackofficeWalletHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/backoffice/wallet/get", _BackofficeWallet_GetWallets0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/wallet/credits/get", _BackofficeWallet_GetWalletCredits0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/wallet/balance-transactions/list", _BackofficeWallet_ListWalletBalanceTransactions0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/wallet/credit-transactions/get", _BackofficeWallet_GetWalletCreditTransactions0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/wallet/update", _BackofficeWallet_UpdateWallet0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/wallet/currencies/add", _BackofficeWallet_AddWalletCurrency0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/wallet/currencies/list", _BackofficeWallet_ListWalletCurrencies0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/wallet/currencies/update", _BackofficeWallet_UpdateWalletCurrency0_HTTP_Handler(srv))
}

func _BackofficeWallet_GetWallets0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWalletsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletGetWallets)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWallets(ctx, req.(*GetWalletsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWalletsResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeWallet_GetWalletCredits0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWalletCreditsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletGetWalletCredits)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWalletCredits(ctx, req.(*GetWalletCreditsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWalletCreditsResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeWallet_ListWalletBalanceTransactions0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListWalletBalanceTransactionsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletListWalletBalanceTransactions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListWalletBalanceTransactions(ctx, req.(*ListWalletBalanceTransactionsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListWalletBalanceTransactionsResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeWallet_GetWalletCreditTransactions0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWalletCreditTransactionsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletGetWalletCreditTransactions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWalletCreditTransactions(ctx, req.(*GetWalletCreditTransactionsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWalletCreditTransactionsResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeWallet_UpdateWallet0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateWalletRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletUpdateWallet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateWallet(ctx, req.(*UpdateWalletRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateWalletResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeWallet_AddWalletCurrency0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddWalletCurrencyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletAddWalletCurrency)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddWalletCurrency(ctx, req.(*AddWalletCurrencyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddWalletCurrencyResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeWallet_ListWalletCurrencies0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListWalletCurrenciesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletListWalletCurrencies)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListWalletCurrencies(ctx, req.(*ListWalletCurrenciesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListWalletCurrenciesResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeWallet_UpdateWalletCurrency0_HTTP_Handler(srv BackofficeWalletHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateWalletCurrencyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeWalletUpdateWalletCurrency)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateWalletCurrency(ctx, req.(*UpdateWalletCurrencyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateWalletCurrencyResponse)
		return ctx.Result(200, reply)
	}
}

type BackofficeWalletHTTPClient interface {
	AddWalletCurrency(ctx context.Context, req *AddWalletCurrencyRequest, opts ...http.CallOption) (rsp *AddWalletCurrencyResponse, err error)
	GetWalletCreditTransactions(ctx context.Context, req *GetWalletCreditTransactionsRequest, opts ...http.CallOption) (rsp *GetWalletCreditTransactionsResponse, err error)
	GetWalletCredits(ctx context.Context, req *GetWalletCreditsRequest, opts ...http.CallOption) (rsp *GetWalletCreditsResponse, err error)
	GetWallets(ctx context.Context, req *GetWalletsRequest, opts ...http.CallOption) (rsp *GetWalletsResponse, err error)
	ListWalletBalanceTransactions(ctx context.Context, req *ListWalletBalanceTransactionsRequest, opts ...http.CallOption) (rsp *ListWalletBalanceTransactionsResponse, err error)
	ListWalletCurrencies(ctx context.Context, req *ListWalletCurrenciesRequest, opts ...http.CallOption) (rsp *ListWalletCurrenciesResponse, err error)
	UpdateWallet(ctx context.Context, req *UpdateWalletRequest, opts ...http.CallOption) (rsp *UpdateWalletResponse, err error)
	UpdateWalletCurrency(ctx context.Context, req *UpdateWalletCurrencyRequest, opts ...http.CallOption) (rsp *UpdateWalletCurrencyResponse, err error)
}

type BackofficeWalletHTTPClientImpl struct {
	cc *http.Client
}

func NewBackofficeWalletHTTPClient(client *http.Client) BackofficeWalletHTTPClient {
	return &BackofficeWalletHTTPClientImpl{client}
}

func (c *BackofficeWalletHTTPClientImpl) AddWalletCurrency(ctx context.Context, in *AddWalletCurrencyRequest, opts ...http.CallOption) (*AddWalletCurrencyResponse, error) {
	var out AddWalletCurrencyResponse
	pattern := "/v1/backoffice/wallet/currencies/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletAddWalletCurrency))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeWalletHTTPClientImpl) GetWalletCreditTransactions(ctx context.Context, in *GetWalletCreditTransactionsRequest, opts ...http.CallOption) (*GetWalletCreditTransactionsResponse, error) {
	var out GetWalletCreditTransactionsResponse
	pattern := "/v1/backoffice/wallet/credit-transactions/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletGetWalletCreditTransactions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeWalletHTTPClientImpl) GetWalletCredits(ctx context.Context, in *GetWalletCreditsRequest, opts ...http.CallOption) (*GetWalletCreditsResponse, error) {
	var out GetWalletCreditsResponse
	pattern := "/v1/backoffice/wallet/credits/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletGetWalletCredits))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeWalletHTTPClientImpl) GetWallets(ctx context.Context, in *GetWalletsRequest, opts ...http.CallOption) (*GetWalletsResponse, error) {
	var out GetWalletsResponse
	pattern := "/v1/backoffice/wallet/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletGetWallets))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeWalletHTTPClientImpl) ListWalletBalanceTransactions(ctx context.Context, in *ListWalletBalanceTransactionsRequest, opts ...http.CallOption) (*ListWalletBalanceTransactionsResponse, error) {
	var out ListWalletBalanceTransactionsResponse
	pattern := "/v1/backoffice/wallet/balance-transactions/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletListWalletBalanceTransactions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeWalletHTTPClientImpl) ListWalletCurrencies(ctx context.Context, in *ListWalletCurrenciesRequest, opts ...http.CallOption) (*ListWalletCurrenciesResponse, error) {
	var out ListWalletCurrenciesResponse
	pattern := "/v1/backoffice/wallet/currencies/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletListWalletCurrencies))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeWalletHTTPClientImpl) UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...http.CallOption) (*UpdateWalletResponse, error) {
	var out UpdateWalletResponse
	pattern := "/v1/backoffice/wallet/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletUpdateWallet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeWalletHTTPClientImpl) UpdateWalletCurrency(ctx context.Context, in *UpdateWalletCurrencyRequest, opts ...http.CallOption) (*UpdateWalletCurrencyResponse, error) {
	var out UpdateWalletCurrencyResponse
	pattern := "/v1/backoffice/wallet/currencies/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeWalletUpdateWalletCurrency))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

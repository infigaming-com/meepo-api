// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.25.6
// source: backoffice/service/v1/backoffice_bcpay.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/infigaming-com/meepo-api/bcpay/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBackofficeBcpayAddBankAccount = "/api.backoffice.service.v1.BackofficeBcpay/AddBankAccount"
const OperationBackofficeBcpayAuditTransaction = "/api.backoffice.service.v1.BackofficeBcpay/AuditTransaction"
const OperationBackofficeBcpayBankAccountList = "/api.backoffice.service.v1.BackofficeBcpay/BankAccountList"
const OperationBackofficeBcpayCreateMerchant = "/api.backoffice.service.v1.BackofficeBcpay/CreateMerchant"
const OperationBackofficeBcpayTransactionList = "/api.backoffice.service.v1.BackofficeBcpay/TransactionList"
const OperationBackofficeBcpayUpdateBankAccount = "/api.backoffice.service.v1.BackofficeBcpay/UpdateBankAccount"

type BackofficeBcpayHTTPServer interface {
	AddBankAccount(context.Context, *v1.AddBankAccountRequest) (*v1.AddBankAccountResponse, error)
	AuditTransaction(context.Context, *v1.AuditTransactionRequest) (*v1.AuditTransactionResponse, error)
	BankAccountList(context.Context, *v1.BankAccountListRequest) (*v1.BankAccountListResponse, error)
	CreateMerchant(context.Context, *v1.CreateMerchantRequest) (*v1.CreateMerchantResponse, error)
	TransactionList(context.Context, *v1.TransactionListRequest) (*v1.TransactionListResponse, error)
	UpdateBankAccount(context.Context, *v1.UpdateBankAccountRequest) (*v1.UpdateBankAccountResponse, error)
}

func RegisterBackofficeBcpayHTTPServer(s *http.Server, srv BackofficeBcpayHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/backoffice/bcpay/merchant/add", _BackofficeBcpay_CreateMerchant0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/bcpay/bankaccount/list", _BackofficeBcpay_BankAccountList0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/bcpay/bankaccount/add", _BackofficeBcpay_AddBankAccount0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/bcpay/bankaccount/update", _BackofficeBcpay_UpdateBankAccount0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/bcpay/transaction/list", _BackofficeBcpay_TransactionList0_HTTP_Handler(srv))
	r.POST("/v1/backoffice/bcpay/transaction/add", _BackofficeBcpay_AuditTransaction0_HTTP_Handler(srv))
}

func _BackofficeBcpay_CreateMerchant0_HTTP_Handler(srv BackofficeBcpayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateMerchantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeBcpayCreateMerchant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMerchant(ctx, req.(*v1.CreateMerchantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.CreateMerchantResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeBcpay_BankAccountList0_HTTP_Handler(srv BackofficeBcpayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.BankAccountListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeBcpayBankAccountList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BankAccountList(ctx, req.(*v1.BankAccountListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.BankAccountListResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeBcpay_AddBankAccount0_HTTP_Handler(srv BackofficeBcpayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.AddBankAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeBcpayAddBankAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddBankAccount(ctx, req.(*v1.AddBankAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.AddBankAccountResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeBcpay_UpdateBankAccount0_HTTP_Handler(srv BackofficeBcpayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateBankAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeBcpayUpdateBankAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBankAccount(ctx, req.(*v1.UpdateBankAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.UpdateBankAccountResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeBcpay_TransactionList0_HTTP_Handler(srv BackofficeBcpayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.TransactionListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeBcpayTransactionList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TransactionList(ctx, req.(*v1.TransactionListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.TransactionListResponse)
		return ctx.Result(200, reply)
	}
}

func _BackofficeBcpay_AuditTransaction0_HTTP_Handler(srv BackofficeBcpayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.AuditTransactionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeBcpayAuditTransaction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuditTransaction(ctx, req.(*v1.AuditTransactionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.AuditTransactionResponse)
		return ctx.Result(200, reply)
	}
}

type BackofficeBcpayHTTPClient interface {
	AddBankAccount(ctx context.Context, req *v1.AddBankAccountRequest, opts ...http.CallOption) (rsp *v1.AddBankAccountResponse, err error)
	AuditTransaction(ctx context.Context, req *v1.AuditTransactionRequest, opts ...http.CallOption) (rsp *v1.AuditTransactionResponse, err error)
	BankAccountList(ctx context.Context, req *v1.BankAccountListRequest, opts ...http.CallOption) (rsp *v1.BankAccountListResponse, err error)
	CreateMerchant(ctx context.Context, req *v1.CreateMerchantRequest, opts ...http.CallOption) (rsp *v1.CreateMerchantResponse, err error)
	TransactionList(ctx context.Context, req *v1.TransactionListRequest, opts ...http.CallOption) (rsp *v1.TransactionListResponse, err error)
	UpdateBankAccount(ctx context.Context, req *v1.UpdateBankAccountRequest, opts ...http.CallOption) (rsp *v1.UpdateBankAccountResponse, err error)
}

type BackofficeBcpayHTTPClientImpl struct {
	cc *http.Client
}

func NewBackofficeBcpayHTTPClient(client *http.Client) BackofficeBcpayHTTPClient {
	return &BackofficeBcpayHTTPClientImpl{client}
}

func (c *BackofficeBcpayHTTPClientImpl) AddBankAccount(ctx context.Context, in *v1.AddBankAccountRequest, opts ...http.CallOption) (*v1.AddBankAccountResponse, error) {
	var out v1.AddBankAccountResponse
	pattern := "/v1/backoffice/bcpay/bankaccount/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeBcpayAddBankAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeBcpayHTTPClientImpl) AuditTransaction(ctx context.Context, in *v1.AuditTransactionRequest, opts ...http.CallOption) (*v1.AuditTransactionResponse, error) {
	var out v1.AuditTransactionResponse
	pattern := "/v1/backoffice/bcpay/transaction/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeBcpayAuditTransaction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeBcpayHTTPClientImpl) BankAccountList(ctx context.Context, in *v1.BankAccountListRequest, opts ...http.CallOption) (*v1.BankAccountListResponse, error) {
	var out v1.BankAccountListResponse
	pattern := "/v1/backoffice/bcpay/bankaccount/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeBcpayBankAccountList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeBcpayHTTPClientImpl) CreateMerchant(ctx context.Context, in *v1.CreateMerchantRequest, opts ...http.CallOption) (*v1.CreateMerchantResponse, error) {
	var out v1.CreateMerchantResponse
	pattern := "/v1/backoffice/bcpay/merchant/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeBcpayCreateMerchant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeBcpayHTTPClientImpl) TransactionList(ctx context.Context, in *v1.TransactionListRequest, opts ...http.CallOption) (*v1.TransactionListResponse, error) {
	var out v1.TransactionListResponse
	pattern := "/v1/backoffice/bcpay/transaction/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeBcpayTransactionList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BackofficeBcpayHTTPClientImpl) UpdateBankAccount(ctx context.Context, in *v1.UpdateBankAccountRequest, opts ...http.CallOption) (*v1.UpdateBankAccountResponse, error) {
	var out v1.UpdateBankAccountResponse
	pattern := "/v1/backoffice/bcpay/bankaccount/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeBcpayUpdateBankAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

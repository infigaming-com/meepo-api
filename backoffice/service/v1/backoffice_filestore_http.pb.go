// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.25.6
// source: backoffice/service/v1/backoffice_filestore.proto

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

const OperationBackofficeFileStoreUploadOperatorStaticFile = "/api.backoffice.service.v1.BackofficeFileStore/UploadOperatorStaticFile"

type BackofficeFileStoreHTTPServer interface {
	UploadOperatorStaticFile(context.Context, *UploadOperatorStaticFileRequest) (*UploadOperatorStaticFileResponse, error)
}

func RegisterBackofficeFileStoreHTTPServer(s *http.Server, srv BackofficeFileStoreHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/backoffice/filestore/operator-static-files/upload", _BackofficeFileStore_UploadOperatorStaticFile0_HTTP_Handler(srv))
}

func _BackofficeFileStore_UploadOperatorStaticFile0_HTTP_Handler(srv BackofficeFileStoreHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadOperatorStaticFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBackofficeFileStoreUploadOperatorStaticFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadOperatorStaticFile(ctx, req.(*UploadOperatorStaticFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadOperatorStaticFileResponse)
		return ctx.Result(200, reply)
	}
}

type BackofficeFileStoreHTTPClient interface {
	UploadOperatorStaticFile(ctx context.Context, req *UploadOperatorStaticFileRequest, opts ...http.CallOption) (rsp *UploadOperatorStaticFileResponse, err error)
}

type BackofficeFileStoreHTTPClientImpl struct {
	cc *http.Client
}

func NewBackofficeFileStoreHTTPClient(client *http.Client) BackofficeFileStoreHTTPClient {
	return &BackofficeFileStoreHTTPClientImpl{client}
}

func (c *BackofficeFileStoreHTTPClientImpl) UploadOperatorStaticFile(ctx context.Context, in *UploadOperatorStaticFileRequest, opts ...http.CallOption) (*UploadOperatorStaticFileResponse, error) {
	var out UploadOperatorStaticFileResponse
	pattern := "/v1/backoffice/filestore/operator-static-files/upload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBackofficeFileStoreUploadOperatorStaticFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

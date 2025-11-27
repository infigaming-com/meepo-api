package middleware

import (
	"context"
	"slices"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/infigaming-com/meepo-api/pkg/util"

	mctx "github.com/infigaming-com/meepo-api/pkg/context"
)

func OperatorStatusCheck(pathIncluder func(string) bool) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				if !pathIncluder(r.URL.Path) {
					return handler(ctx, req)
				}
				requestInfo, ok := mctx.GetRequestInfo(ctx)
				if !ok {
					return nil, errors.New(403, "FORBIDDEN", "invalid system operatorId")
				}
				operatorInfo, ok := mctx.GetOperatorInfo(ctx)
				if !ok {
					return nil, errors.New(403, "FORBIDDEN", "GetOperatorInfo failed")
				}

				if operatorInfo.Status == util.OperatorStatusClosed {
					return nil, errors.New(403, "FORBIDDEN", "website close")
				}

				if operatorInfo.Status == util.OperatorStatusMaintain && !slices.Contains(operatorInfo.StatusLaunchWhitelist, requestInfo.ClientIP) {
					return nil, errors.New(403, "FORBIDDEN", "IP is not allow to access").WithMetadata(map[string]string{
						"status":  operatorInfo.Status,
						"endtime": strconv.FormatInt(operatorInfo.StatusEndTime, 10),
					})
				}
			}
			return handler(ctx, req)
		}
	}
}

func OperatorStatusCheckWithExcludePaths(excludePaths []string) middleware.Middleware {
	return OperatorStatusCheck(func(path string) bool {
		return !slices.Contains(excludePaths, path)
	})
}

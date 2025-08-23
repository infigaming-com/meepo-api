package middleware

import (
	"context"
	"slices"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/infigaming-com/meepo-api/pkg/util"

	mctx "github.com/infigaming-com/meepo-api/pkg/context"
)

func OperatorStatusCheck() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			requestInfo, ok := mctx.GetRequestInfo(ctx)
			if !ok {
				return nil, errors.New(403, "FORBIDDEN", "invalid system operatorId")
			}
			operatorInfo, ok := mctx.GetOperatorInfo(ctx)
			if !ok {
				return nil, errors.New(403, "FORBIDDEN", "GetOperatorInfo failed")
			}

			if operatorInfo.Status == util.OperatorStatusPending && !slices.Contains(operatorInfo.StatusLaunchWhitelist, requestInfo.ClientIP) {
				return nil, errors.New(403, "FORBIDDEN", "IP is not allow to access")
			}
			return handler(ctx, req)
		}
	}
}

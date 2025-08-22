package middleware

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/infigaming-com/meepo-api/pkg/util"

	mctx "github.com/infigaming-com/meepo-api/pkg/context"
)

func OperatorStatusCheck() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			requestInfo, ok := mctx.GetRequestInfo(ctx)
			if !ok {
				return nil, errors.New("GetRequestInfo failed")
			}
			operatorInfo, ok := mctx.GetOperatorInfo(ctx)
			if !ok {
				return nil, errors.New("GetOperatorInfo failed")
			}

			if operatorInfo.Status == util.OperatorStatusPending && slices.Contains(operatorInfo.StatusLaunchWhitelist, requestInfo.ClientIP) {
				return nil, fmt.Errorf("IP is not allow to access (%s)", requestInfo.ClientIP)
			}
			return handler(ctx, req)
		}
	}
}

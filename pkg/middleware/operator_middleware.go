package middleware

import (
	"context"
	"slices"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	user "github.com/infigaming-com/meepo-api/user/service/v1"
)

// OperatorIdMiddlewareWithPathIncluder is a middleware that extract Origin and
// get the operatorId with Origin from redis, using a custom function to determine
// which paths should be processed
func OperatorIdMiddlewareWithPathIncluder(pathIncluder func(string) bool, userClient user.UserClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				if !pathIncluder(r.URL.Path) {
					return handler(ctx, req)
				}
				origin := r.Header.Get("Origin")
				if origin == "" {
					return nil, errors.New(400, "BAD_REQUEST", "missing origin header")
				}
				// temporary use map to store origin and operatorId
				resp, err := userClient.GetOperatorIdsByOrigin(ctx, &user.GetOperatorIdsByOriginRequest{
					Origin: origin,
				})
				if err != nil {
					if user.IsOperatorIdsNotFoundByOrigin(err) {
						return nil, errors.New(400, "BAD_REQUEST", "operatorIds not found by origin")
					} else {
						return nil, errors.New(400, "BAD_REQUEST", "cannot get operatorIds by origin")
					}
				}
				operatorIds := mctx.OperatorIdsFromOperatorContext(resp.OperatorContext)
				operatorIds.RealOperatorId, operatorIds.OperatorType = operatorIds.GetRealOperatorIdAndType()
				ctx = mctx.WithOperatorIds(ctx, operatorIds)
			}
			return handler(ctx, req)
		}
	}
}

// OperatorIdMiddlewareWithExcludePaths is a middleware that extract Origin and
// get the operatorId with Origin from redis, excluding specific paths from processing
func OperatorIdMiddlewareWithExcludePaths(excludePaths []string, userClient user.UserClient) middleware.Middleware {
	return OperatorIdMiddlewareWithPathIncluder(func(path string) bool {
		return !slices.Contains(excludePaths, path)
	}, userClient)
}

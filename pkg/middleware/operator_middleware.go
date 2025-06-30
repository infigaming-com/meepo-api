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

// OperatorIdMiddleware is a middleware that extract Origin and
// get the operatorId with Origin from redis
func OperatorIdMiddleware(path []string, userClient user.UserClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				if len(path) > 0 && !slices.Contains(path, r.URL.Path) {
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
				ctx = mctx.WithOperators(ctx, mctx.Operators{
					OperatorId:         resp.OperatorContext.OperatorId,
					CompanyOperatorId:  resp.OperatorContext.CompanyOperatorId,
					RetailerOperatorId: resp.OperatorContext.RetailerOperatorId,
					SystemOperatorId:   resp.OperatorContext.SystemOperatorId,
				})
			}
			return handler(ctx, req)
		}
	}
}

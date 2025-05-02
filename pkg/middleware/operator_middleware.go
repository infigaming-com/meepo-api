package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	operator "github.com/infigaming-com/meepo-api/operator/service/v1"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
)

type operatorKey struct{}

var (
	OperatorKey           = operatorKey{}
	OriginToOperatorIdMap = map[string]int64{
		"https://dev.mini.bet": 1234567890,
	}
)

// OperatorIdMiddleware is a middleware that extract Origin and
// get the operatorId with Origin from redis
func OperatorIdMiddleware(lg log.Logger, operator operator.OperatorClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				origin := r.Header.Get("Origin")
				if origin == "" {
					return nil, errors.New(400, "BAD_REQUEST", "missing origin header")
				}
				// temporary use map to store origin and operatorId
				operatorId, ok := OriginToOperatorIdMap[origin]
				if !ok {
					//return nil, errors.New(400, "BAD_REQUEST", "invalid origin")
					operatorId = 1234567890 // temporary use 1234567890 as default operatorId
				}
				ctx = mctx.WithOperatorId(ctx, operatorId)
			}
			return handler(ctx, req)
		}
	}
}

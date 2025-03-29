package middleware

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// JwtExtractor is a middleware that extracts the JWT token from the Authorization header
// and adds it to the context.
func JwtExtractor() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			authorization := ""
			if tr, ok := transport.FromServerContext(ctx); ok {
				switch tr := tr.(type) {
				case *http.Transport:
					authorization = tr.RequestHeader().Get("Authorization")
				case *grpc.Transport:
					authorization = tr.RequestHeader().Get("Authorization")
				}
			}
			if authorization != "" && strings.HasPrefix(authorization, "Bearer ") {
				token := strings.TrimPrefix(authorization, "Bearer ")
				if token != "" {
					ctx = context.WithValue(ctx, "jwt_token", token)
				}
			}
			return handler(ctx, req)
		}
	}
}

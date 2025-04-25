package middleware

import (
	"context"
	"slices"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	"github.com/infigaming-com/meepo-api/pkg/jwt"
	user "github.com/infigaming-com/meepo-api/user/service/v1"
)

func AuthMiddleware(authPaths []string, secret string, uc user.UserClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New(400, "BAD_REQUEST", "no transport")
			}
			httpTr, ok := tr.(*http.Transport)
			if !ok {
				return nil, errors.New(400, "BAD_REQUEST", "not http transport")
			}

			if !slices.Contains(authPaths, httpTr.Request().URL.Path) {
				return handler(ctx, req)
			}

			authorization := httpTr.Request().Header.Get("Authorization")
			if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
				return nil, errors.New(401, "UNAUTHORIZED", "no authorization")
			}

			token, claims, err := jwt.ParseToken(authorization, secret)
			if err != nil {
				return nil, errors.New(401, "UNAUTHORIZED", "invalid token")
			}

			revoked, err := uc.IsTokenRevoked(ctx, &user.IsTokenRevokedRequest{
				Token: token,
			})
			if err != nil || revoked.Revoked {
				return nil, errors.New(401, "UNAUTHORIZED", "invalid token")
			}

			ctx = mctx.WithUserInfo(ctx, claims.UserInfo)

			return handler(ctx, req)
		}
	}
}

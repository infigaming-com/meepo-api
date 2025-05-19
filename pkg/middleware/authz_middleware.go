package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	user "github.com/infigaming-com/meepo-api/user/service/v1"
)

func AuthzMiddleware(secret string, uc user.UserClient) middleware.Middleware {
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
			path := httpTr.Request().URL.Path

			userInfo, ok := mctx.UserInfo(ctx)
			if !ok {
				return nil, errors.New(401, "UNAUTHORIZED", "no user info")
			}
			allowed, err := uc.CheckPermission(ctx, &user.CheckPermissionRequest{
				Path:       path,
				RoleId:     userInfo.RoleId,
				OperatorId: userInfo.OperatorId,
			})
			if err != nil || !allowed.Allowed {
				return nil, errors.New(401, "UNAUTHORIZED", "no permission")
			}

			return handler(ctx, req)
		}
	}
}

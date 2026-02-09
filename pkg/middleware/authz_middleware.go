package middleware

import (
	"context"
	"slices"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	user "github.com/infigaming-com/meepo-api/user/service/v1"
)

func AuthzMiddleware(paths []string, secret string, uc user.UserClient) middleware.Middleware {
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
			if !slices.Contains(paths, path) {
				return handler(ctx, req)
			}

			userInfo, ok := mctx.UserInfo(ctx)
			if !ok {
				return nil, errors.New(401, "UNAUTHORIZED", "no user info")
			}
			operatorId := userInfo.OperatorId
			if userInfo.RealOperatorId != 0 {
				operatorId = userInfo.RealOperatorId
			}
			allowed, err := uc.CheckPermission(ctx, &user.CheckPermissionRequest{
				Path:       path,
				RoleId:     userInfo.RoleId,
				OperatorId: operatorId,
			})
			if err != nil || !allowed.Allowed {
				return nil, errors.New(401, "UNAUTHORIZED", "no permission")
			}

			return handler(ctx, req)
		}
	}
}

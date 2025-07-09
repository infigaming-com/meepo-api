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
			if authorization == "" {
				return nil, errors.New(401, "UNAUTHORIZED", "no authorization")
			}
			if !strings.HasPrefix(authorization, "Bearer ") {
				return nil, errors.New(401, "UNAUTHORIZED", "invalid authorization")
			}

			token, claims, err := jwt.ParseToken(authorization, secret)
			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					return nil, errors.New(401, "TOKEN_EXPIRED", "token expired")
				}
				return nil, errors.New(401, "UNAUTHORIZED", "invalid token")
			}

			operatorIds, ok := mctx.GetOperatorIds(ctx)
			// If the operatorId in the context does not exists, do not check the operatorId in the token
			// If the operatorId in the context exists, check if it is the same as the operatorId in the token
			if ok {
				if claims.UserInfo.OperatorId != operatorIds.OperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid operatorId")
				}
				if claims.UserInfo.CompanyOperatorId != operatorIds.CompanyOperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid company operatorId")
				}
				if claims.UserInfo.RetailerOperatorId != operatorIds.RetailerOperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid retailer operatorId")
				}
				if claims.UserInfo.SystemOperatorId != operatorIds.SystemOperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid system operatorId")
				}
			}

			revoked, err := uc.IsTokenRevoked(ctx, &user.IsTokenRevokedRequest{
				Token: token,
			})
			if err != nil || revoked.Revoked {
				return nil, errors.New(401, "UNAUTHORIZED", "invalid token")
			}

			userOperatorIds := &mctx.OperatorIds{
				OperatorId:         claims.UserInfo.OperatorId,
				CompanyOperatorId:  claims.UserInfo.CompanyOperatorId,
				RetailerOperatorId: claims.UserInfo.RetailerOperatorId,
				SystemOperatorId:   claims.UserInfo.SystemOperatorId,
			}
			claims.UserInfo.RealOperatorId, claims.UserInfo.OperatorType = userOperatorIds.GetRealOperatorIdAndType()

			ctx = mctx.WithUserInfo(ctx, claims.UserInfo)

			return handler(ctx, req)
		}
	}
}

package middleware

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	"github.com/infigaming-com/meepo-api/pkg/jwt"
	user "github.com/infigaming-com/meepo-api/user/service/v1"
)

func AuthMiddlewareWithPathIncluder(pathIncluder func(string) bool, secret string, uc user.UserClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New(400, "BAD_REQUEST", "no transport")
			}
			httpTr, ok := tr.(*http.Transport)
			if !ok {
				return nil, errors.New(400, "BAD_REQUEST", "not http transport")
			}

			if !pathIncluder(httpTr.Request().URL.Path) {
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

			operatorContext, ok := mctx.GetOperatorContext(ctx)
			// If the operatorId in the context does not exists, do not check the operatorId in the token
			// If the operatorId in the context exists, check if it is the same as the operatorId in the token
			if ok {
				if claims.UserInfo.OperatorId != operatorContext.OperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid operatorId")
				}
				if claims.UserInfo.CompanyOperatorId != operatorContext.CompanyOperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid company operatorId")
				}
				if claims.UserInfo.RetailerOperatorId != operatorContext.RetailerOperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid retailer operatorId")
				}
				if claims.UserInfo.SystemOperatorId != operatorContext.SystemOperatorId {
					return nil, errors.New(401, "UNAUTHORIZED", "invalid system operatorId")
				}
			}

			if claims.UserInfo.NeedResetPassword {
				return nil, errors.New(403, "NEED_RESET_PASSWORD", "password reset required before access")
			}

			operatorInfo, ok := mctx.GetOperatorInfo(ctx)
			if !ok {
				return nil, errors.New(401, "UNAUTHORIZED", "get operator info failed")
			}

			if operatorInfo.Config != nil &&
				operatorInfo.Config.AccountSettings != nil &&
				operatorInfo.Config.AccountSettings.PasswordExpiryDays != 0 &&
				claims.UserInfo.PasswordResetAt != 0 { // for user with google login, there won't be password
				passwordResetAt := time.UnixMilli(claims.UserInfo.PasswordResetAt)
				passwordExpiredAt := passwordResetAt.AddDate(0, 0, int(operatorInfo.Config.AccountSettings.PasswordExpiryDays))
				if time.Now().After(passwordExpiredAt) {
					return nil, errors.New(403, "PASSWORD_EXPIRED", "password expired")
				}
			}

			revoked, err := uc.IsTokenRevoked(ctx, &user.IsTokenRevokedRequest{
				Token: token,
			})
			if err != nil || revoked.Revoked {
				return nil, errors.New(401, "UNAUTHORIZED", "invalid token")
			}

			claims.UserInfo.RealOperatorId = operatorContext.RealOperatorId
			claims.UserInfo.OperatorType = operatorContext.OperatorType
			ctx = mctx.WithUserInfo(ctx, claims.UserInfo)

			return handler(ctx, req)
		}
	}
}

func AuthMiddlewareWithExcludePaths(excludePaths []string, secret string, uc user.UserClient) middleware.Middleware {
	return AuthMiddlewareWithPathIncluder(func(path string) bool {
		return !slices.Contains(excludePaths, path)
	}, secret, uc)
}

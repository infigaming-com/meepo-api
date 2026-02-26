package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	"github.com/infigaming-com/meepo-api/pkg/util"
)

// SessionActivityMiddleware tracks session activity for players (RoleId == 0).
// The trackFn is called with extracted user and request information for every
// authenticated player request. Tracking logic (deduplication, caching) is
// handled by the caller-provided function.
func SessionActivityMiddleware(
	trackFn func(ctx context.Context, userId, operatorId int64, ip, userAgent, country string),
) middleware.Middleware {
	return func(next middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if userInfo, ok := mctx.UserInfo(ctx); ok && userInfo.RoleId == util.PlayerRoleId {
				if reqInfo, ok := mctx.GetRequestInfo(ctx); ok {
					trackFn(ctx, userInfo.UserId, userInfo.RealOperatorId,
						reqInfo.ClientIP, reqInfo.UserAgent, reqInfo.Country)
				}
			}
			return next(ctx, req)
		}
	}
}

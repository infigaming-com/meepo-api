package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	"github.com/infigaming-com/meepo-api/pkg/util"
)

// ClientSourceHeader is the HTTP header name for client source identification.
const ClientSourceHeader = "X-Client-Source"

// Client source values.
const (
	ClientSourcePWA = "pwa"
)

// SessionActivityRequest carries enriched session info for the tracker.
type SessionActivityRequest struct {
	UserID             int64
	RealOperatorID     int64
	OperatorID         int64
	CompanyOperatorID  int64
	RetailerOperatorID int64
	SystemOperatorID   int64
	OperatorType       string
	IP                 string
	UserAgent          string
	Country            string
	ClientSource       string
}

// SessionActivityMiddleware tracks session activity for players (RoleId == 0).
// The trackFn is called with extracted user and request information for every
// authenticated player request. Tracking logic (deduplication, caching) is
// handled by the caller-provided function.
func SessionActivityMiddleware(
	trackFn func(ctx context.Context, req *SessionActivityRequest),
) middleware.Middleware {
	return func(next middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if userInfo, ok := mctx.UserInfo(ctx); ok && userInfo.RoleId == util.PlayerRoleId {
				if reqInfo, ok := mctx.GetRequestInfo(ctx); ok {
					trackFn(ctx, &SessionActivityRequest{
						UserID:             userInfo.UserId,
						RealOperatorID:     userInfo.RealOperatorId,
						OperatorID:         userInfo.OperatorId,
						CompanyOperatorID:  userInfo.CompanyOperatorId,
						RetailerOperatorID: userInfo.RetailerOperatorId,
						SystemOperatorID:   userInfo.SystemOperatorId,
						OperatorType:       userInfo.OperatorType,
						IP:                 reqInfo.ClientIP,
						UserAgent:          reqInfo.UserAgent,
						Country:            reqInfo.Country,
						ClientSource:       reqInfo.ClientSource,
					})
				}
			}
			return next(ctx, req)
		}
	}
}

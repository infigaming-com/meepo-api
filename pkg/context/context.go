package context

import (
	"context"

	"github.com/infigaming-com/meepo-api/pkg/jwt"
)

type contextKey string

type RequestInfo struct {
	Method     string
	Path       string
	Host       string
	RemoteAddr string
	Country    string
	Origin     string
	Referer    string
	UserAgent  string
	ClientIP   string
}

type OperatorIds struct {
	OperatorId         int64
	CompanyOperatorId  int64
	RetailerOperatorId int64
	SystemOperatorId   int64
}

func WithValue[T any](ctx context.Context, key string, value T) context.Context {
	return context.WithValue(ctx, contextKey(key), value)
}

func Value[T any](ctx context.Context, key string) (T, bool) {
	var zero T
	if v := ctx.Value(contextKey(key)); v != nil {
		if s, ok := v.(T); ok {
			return s, true
		}
	}
	return zero, false
}

func WithUserInfo(ctx context.Context, userInfo jwt.UserInfo) context.Context {
	return WithValue(ctx, "userInfo", userInfo)
}

func UserInfo(ctx context.Context) (jwt.UserInfo, bool) {
	return Value[jwt.UserInfo](ctx, "userInfo")
}

// deprecated
func WithOperatorId(ctx context.Context, operatorId int64) context.Context {
	return WithValue(ctx, "operatorId", operatorId)
}

// deprecated
func OperatorId(ctx context.Context) (int64, bool) {
	return Value[int64](ctx, "operatorId")
}

func WithOperatorIds(ctx context.Context, operatorIds OperatorIds) context.Context {
	return WithValue(ctx, "operatorIds", operatorIds)
}

func GetOperatorIds(ctx context.Context) (OperatorIds, bool) {
	return Value[OperatorIds](ctx, "operatorIds")
}

func WithRequestInfo(ctx context.Context, requestInfo RequestInfo) context.Context {
	return WithValue(ctx, "requestInfo", requestInfo)
}

func GetRequestInfo(ctx context.Context) (RequestInfo, bool) {
	return Value[RequestInfo](ctx, "requestInfo")
}

func GetActualOperatorIdAndType(ctx context.Context) (int64, string, bool) {
	if operatorIds, ok := GetOperatorIds(ctx); ok {
		if operatorIds.OperatorId != 0 {
			return operatorIds.OperatorId, "operator", true
		} else if operatorIds.CompanyOperatorId != 0 {
			return operatorIds.CompanyOperatorId, "company", true
		} else if operatorIds.RetailerOperatorId != 0 {
			return operatorIds.RetailerOperatorId, "retailer", true
		} else {
			return operatorIds.SystemOperatorId, "system", true
		}
	}
	return 0, "", false
}

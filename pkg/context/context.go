package context

import (
	"context"

	"github.com/infigaming-com/meepo-api/pkg/jwt"
)

type contextKey string

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

func WithOperatorId(ctx context.Context, operatorId int64) context.Context {
	return WithValue(ctx, "operatorId", operatorId)
}

func OperatorId(ctx context.Context) (int64, bool) {
	return Value[int64](ctx, "operatorId")
}

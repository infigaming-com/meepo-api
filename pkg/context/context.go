package context

import (
	"context"

	"github.com/infigaming-com/meepo-api/common"
	"github.com/infigaming-com/meepo-api/pkg/jwt"
	"github.com/infigaming-com/meepo-api/pkg/util"
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
		actualOperatorId, actualOperatorType := operatorIds.GetActualOperatorIdAndType()
		return actualOperatorId, actualOperatorType, true
	}
	return 0, "", false
}

func (o *OperatorIds) GetActualOperatorIdAndType() (int64, string) {
	if o.OperatorId != 0 {
		// Operator level
		return o.OperatorId, util.OperatorTypeOperator
	} else if o.CompanyOperatorId != 0 {
		// Company level
		return o.CompanyOperatorId, util.OperatorTypeCompany
	} else if o.RetailerOperatorId != 0 {
		// Retailer level
		return o.RetailerOperatorId, util.OperatorTypeRetailer
	} else {
		// System level
		return o.SystemOperatorId, util.OperatorTypeSystem
	}
}

func OperatorIdsFromOperatorContext(oc *common.OperatorContext) OperatorIds {
	return OperatorIds{
		OperatorId:         oc.OperatorId,
		CompanyOperatorId:  oc.CompanyOperatorId,
		RetailerOperatorId: oc.RetailerOperatorId,
		SystemOperatorId:   oc.SystemOperatorId,
	}
}

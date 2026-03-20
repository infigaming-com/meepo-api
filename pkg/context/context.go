package context

import (
	"context"

	"github.com/infigaming-com/meepo-api/common"
	"github.com/infigaming-com/meepo-api/pkg/jwt"
	"github.com/infigaming-com/meepo-api/pkg/util"
)

type contextKey string

type RequestInfo struct {
	Method      string
	Path        string
	Host        string
	RemoteAddr  string
	Country     string
	Origin      string
	Referer     string
	UserAgent   string
	ClientIP    string
	LoginMethod string // from "Login-Method" header
}

type OperatorIds struct {
	OperatorId         int64
	CompanyOperatorId  int64
	RetailerOperatorId int64
	SystemOperatorId   int64
	RealOperatorId     int64
	OperatorType       string
}

type OperatorAccountPasswordSettings struct {
	MinCharacters          int32
	MinUppercaseCharacters int32
	MinLowercaseCharacters int32
	MinSpecialCharacters   int32
	MinDigits              int32
}

type OperatorAccountSecuritySettings struct {
	MaxPasswordRetries    int32
	PasswordExpiryDays    int32
	PasswordHistoryLimits int32
}

type OperatorAccountGameSettings struct {
	NoGameWithoutDeposit bool
	MinGameKYCLevel      int32
}

type OperatorAccountPaymentSettings struct {
	MinDepositKYCLevel  int32
	MinWithdrawKYCLevel int32
}

type OperatorAccountSettings struct {
	PasswordExpiryDays int32
}

type OperatorConfig struct {
	AccountSettings *OperatorAccountSettings
}

type OperatorInfo struct {
	Id                    int64
	OperatorName          string
	ParentOperatorId      int64
	ParentOperatorName    string
	OperatorType          string
	ExternalId            string
	Subdomain             string
	BackofficeSubdomain   string
	BackofficeChildDomain string
	DomainPool            []string
	Enabled               bool
	Mode                  string
	OperatorKey           string
	ReportingCurrency     string
	BackofficeTimezone    string
	SupportedLanguages    []string
	SupportedCurrencies   []string
	Status                string
	IsMaintenance         bool
	StatusStartTime       int64
	StatusEndTime         int64
	OperatorId            int64
	CompanyOperatorId     int64
	CompanyOperatorName   string
	RetailerOperatorId    int64
	RetailerOperatorName  string
	SystemOperatorId      int64
	SystemOperatorName    string
	Config                *OperatorConfig
	MinLaunchBalance      string
	StatusLaunchWhitelist []string

	// System-level maintenance info (populated when system operator is in maintenance)
	SystemStatus                string
	SystemStatusEndTime         int64
	SystemStatusLaunchWhitelist []string
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

func WithOperatorInfo(ctx context.Context, operatorInfo OperatorInfo) context.Context {
	return WithValue(ctx, "operatorInfo", operatorInfo)
}

func GetOperatorInfo(ctx context.Context) (OperatorInfo, bool) {
	return Value[OperatorInfo](ctx, "operatorInfo")
}

func GetOperatorContext(ctx context.Context) (*common.OperatorContext, bool) {
	operatorIds, ok := GetOperatorIds(ctx)
	if !ok {
		return nil, false
	}
	operatorContext := operatorIds.GetOperatorContext()
	return &operatorContext, true
}

func WithRequestInfo(ctx context.Context, requestInfo RequestInfo) context.Context {
	return WithValue(ctx, "requestInfo", requestInfo)
}

func GetRequestInfo(ctx context.Context) (RequestInfo, bool) {
	return Value[RequestInfo](ctx, "requestInfo")
}

// GetLoginAsUserId returns the user ID that the current user is acting as.
// For sub-accounts: returns LoginAsUserId (parent affiliate's userId).
// For everyone else: returns UserId.
func GetLoginAsUserId(userInfo jwt.UserInfo) int64 {
	if userInfo.LoginAsUserId > 0 {
		return userInfo.LoginAsUserId
	}
	return userInfo.UserId
}

// IsSubAccount returns true if the current user is a sub-account of an affiliate.
func IsSubAccount(userInfo jwt.UserInfo) bool {
	return userInfo.LoginAsUserId > 0 && userInfo.LoginAsUserId != userInfo.UserId
}

func (o *OperatorIds) GetRealOperatorIdAndType() (int64, string) {
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

func NewOperatorIds(operatorId, companyOperatorId, retailerOperatorId, systemOperatorId int64) *OperatorIds {
	operatorIds := &OperatorIds{
		OperatorId:         operatorId,
		CompanyOperatorId:  companyOperatorId,
		RetailerOperatorId: retailerOperatorId,
		SystemOperatorId:   systemOperatorId,
	}
	realOperatorId, operatorType := operatorIds.GetRealOperatorIdAndType()
	operatorIds.RealOperatorId = realOperatorId
	operatorIds.OperatorType = operatorType
	return operatorIds
}

func (o *OperatorIds) GetOperatorContext() common.OperatorContext {
	return common.OperatorContext{
		OperatorId:         o.OperatorId,
		CompanyOperatorId:  o.CompanyOperatorId,
		RetailerOperatorId: o.RetailerOperatorId,
		SystemOperatorId:   o.SystemOperatorId,
		RealOperatorId:     o.RealOperatorId,
		OperatorType:       o.OperatorType,
	}
}

func OperatorIdsFromOperatorContext(oc *common.OperatorContext) OperatorIds {
	return OperatorIds{
		OperatorId:         oc.OperatorId,
		CompanyOperatorId:  oc.CompanyOperatorId,
		RetailerOperatorId: oc.RetailerOperatorId,
		SystemOperatorId:   oc.SystemOperatorId,
		RealOperatorId:     oc.RealOperatorId,
		OperatorType:       oc.OperatorType,
	}
}

// IsOperatorIdsInOperatorContext checks if the operatorIds are in the operatorContext.
func IsOperatorIdsInOperatorContext(operatorIds OperatorIds, operatorContext *common.OperatorContext) bool {
	operatorType := operatorContext.OperatorType
	if operatorType == "" {
		contextOperatorIds := OperatorIdsFromOperatorContext(operatorContext)
		_, operatorType = contextOperatorIds.GetRealOperatorIdAndType()
	}
	switch operatorType {
	case util.OperatorTypeOperator:
		return operatorIds.OperatorId == operatorContext.OperatorId && operatorIds.CompanyOperatorId == operatorContext.CompanyOperatorId && operatorIds.RetailerOperatorId == operatorContext.RetailerOperatorId && operatorIds.SystemOperatorId == operatorContext.SystemOperatorId
	case util.OperatorTypeCompany:
		return operatorIds.CompanyOperatorId == operatorContext.CompanyOperatorId && operatorIds.RetailerOperatorId == operatorContext.RetailerOperatorId && operatorIds.SystemOperatorId == operatorContext.SystemOperatorId
	case util.OperatorTypeRetailer:
		return operatorIds.RetailerOperatorId == operatorContext.RetailerOperatorId && operatorIds.SystemOperatorId == operatorContext.SystemOperatorId
	case util.OperatorTypeSystem:
		return operatorIds.SystemOperatorId == operatorContext.SystemOperatorId
	}
	return false
}

func IsTargetOperatorContextInOperatorContext(targetOperatorContext *common.OperatorContext, operatorContext *common.OperatorContext) bool {
	if targetOperatorContext == nil {
		return true
	}
	return IsOperatorIdsInOperatorContext(OperatorIdsFromOperatorContext(targetOperatorContext), operatorContext)
}

func NewOperatorContextWithIds(operatorId, companyOperatorId, retailerOperatorId, systemOperatorId int64) *common.OperatorContext {
	operatorIds := NewOperatorIds(operatorId, companyOperatorId, retailerOperatorId, systemOperatorId)
	operatorContext := operatorIds.GetOperatorContext()
	return &operatorContext
}

func GetParentOperatorIdsSliceFromOperatorContext(operatorContext *common.OperatorContext) []int64 {
	switch operatorContext.OperatorType {
	case util.OperatorTypeOperator:
		return []int64{operatorContext.CompanyOperatorId, operatorContext.RetailerOperatorId, operatorContext.SystemOperatorId}
	case util.OperatorTypeCompany:
		return []int64{operatorContext.RetailerOperatorId, operatorContext.SystemOperatorId}
	case util.OperatorTypeRetailer:
		return []int64{operatorContext.SystemOperatorId}
	case util.OperatorTypeSystem:
		return []int64{}
	}
	return []int64{}
}

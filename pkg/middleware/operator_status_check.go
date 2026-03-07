package middleware

import (
	"context"
	"slices"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
	"github.com/infigaming-com/meepo-api/pkg/util"
)

// OperatorStatusCheck checks operator maintenance status for frontend services.
// It checks system-level maintenance first (higher priority), then operator-level.
func OperatorStatusCheck(pathIncluder func(string) bool) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				if !pathIncluder(r.URL.Path) {
					return handler(ctx, req)
				}
				requestInfo, ok := mctx.GetRequestInfo(ctx)
				if !ok {
					return nil, errors.New(403, "FORBIDDEN", "invalid system operatorId")
				}
				operatorInfo, ok := mctx.GetOperatorInfo(ctx)
				if !ok {
					return nil, errors.New(403, "FORBIDDEN", "GetOperatorInfo failed")
				}

				// 1. Check closed/archived status
				if operatorInfo.Status == util.OperatorStatusClosed || operatorInfo.Status == util.OperatorStatusArchived {
					return nil, errors.New(403, "FORBIDDEN", "website close")
				}

				// 2. Check system-level maintenance FIRST (higher priority)
				if operatorInfo.SystemStatus == util.OperatorStatusMaintain {
					// System maintenance: ONLY system whitelist works
					if !slices.Contains(operatorInfo.SystemStatusLaunchWhitelist, requestInfo.ClientIP) {
						return nil, errors.New(403, "FORBIDDEN", "IP is not allow to access").WithMetadata(map[string]string{
							"status":  util.OperatorStatusMaintain,
							"endtime": strconv.FormatInt(operatorInfo.SystemStatusEndTime, 10),
							"source":  util.OperatorTypeSystem,
						})
					}
					// IP in system whitelist → allow through (skip operator-level check)
					return handler(ctx, req)
				}

				// 3. Check operator-level maintenance
				if operatorInfo.Status == util.OperatorStatusMaintain &&
					!slices.Contains(operatorInfo.StatusLaunchWhitelist, requestInfo.ClientIP) {
					return nil, errors.New(403, "FORBIDDEN", "IP is not allow to access").WithMetadata(map[string]string{
						"status":  operatorInfo.Status,
						"endtime": strconv.FormatInt(operatorInfo.StatusEndTime, 10),
						"source":  util.OperatorTypeOperator,
					})
				}
			}
			return handler(ctx, req)
		}
	}
}

func OperatorStatusCheckWithExcludePaths(excludePaths []string) middleware.Middleware {
	return OperatorStatusCheck(func(path string) bool {
		return !slices.Contains(excludePaths, path)
	})
}

// BackofficeOperatorStatusCheck checks operator maintenance status for backoffice services.
// It allows users at or above the maintenance source level to pass through.
func BackofficeOperatorStatusCheck(pathExcluder func(string) bool) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				// Skip excluded paths (login, register, etc.)
				if pathExcluder(r.URL.Path) {
					return handler(ctx, req)
				}

				operatorInfo, ok := mctx.GetOperatorInfo(ctx)
				if !ok {
					return nil, errors.New(403, "FORBIDDEN", "GetOperatorInfo failed")
				}

				// Closed/Archived → always block
				if operatorInfo.Status == util.OperatorStatusClosed || operatorInfo.Status == util.OperatorStatusArchived {
					return nil, errors.New(403, "FORBIDDEN", "website close")
				}

				// Determine maintenance source and level
				maintenanceLevel := -1 // no maintenance
				maintenanceEndTime := int64(0)
				if operatorInfo.SystemStatus == util.OperatorStatusMaintain {
					maintenanceLevel = operatorTypeToLevel(util.OperatorTypeSystem)
					maintenanceEndTime = operatorInfo.SystemStatusEndTime
				} else if operatorInfo.Status == util.OperatorStatusMaintain {
					maintenanceLevel = operatorTypeToLevel(util.OperatorTypeOperator)
					maintenanceEndTime = operatorInfo.StatusEndTime
				}

				if maintenanceLevel < 0 {
					// No maintenance
					return handler(ctx, req)
				}

				// Get user info (set by AuthMiddleware)
				userInfo, ok := mctx.UserInfo(ctx)
				if !ok {
					// No auth info → block during maintenance
					return nil, errors.New(403, "FORBIDDEN", "site under maintenance").WithMetadata(map[string]string{
						"status":  util.OperatorStatusMaintain,
						"endtime": strconv.FormatInt(maintenanceEndTime, 10),
					})
				}

				// User at maintenance source level or above → allow through
				userLevel := operatorTypeToLevel(userInfo.OperatorType)
				if userLevel >= maintenanceLevel {
					return handler(ctx, req)
				}

				// User below maintenance level → block
				return nil, errors.New(403, "FORBIDDEN", "site under maintenance").WithMetadata(map[string]string{
					"status":  util.OperatorStatusMaintain,
					"endtime": strconv.FormatInt(maintenanceEndTime, 10),
				})
			}
			return handler(ctx, req)
		}
	}
}

// operatorTypeToLevel maps operator type to hierarchy level.
func operatorTypeToLevel(operatorType string) int {
	switch operatorType {
	case util.OperatorTypeSystem:
		return 3
	case util.OperatorTypeRetailer:
		return 2
	case util.OperatorTypeCompany:
		return 1
	case util.OperatorTypeOperator:
		return 0
	default:
		return -1
	}
}

func BackofficeOperatorStatusCheckWithExcludePaths(excludePaths []string) middleware.Middleware {
	return BackofficeOperatorStatusCheck(func(path string) bool {
		return slices.Contains(excludePaths, path)
	})
}

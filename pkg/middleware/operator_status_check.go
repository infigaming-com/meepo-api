package middleware

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func OperatorStatusCheck(getIPs func(ctx context.Context) ([]string, bool)) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				// Get client IP in order of preference
				clientIP := ""
				if cfIP := r.Header.Get("Cf-Connecting-Ip"); cfIP != "" {
					clientIP = cfIP
				} else if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
					clientIP = realIP
				} else if xForwardedFor := r.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
					// Take the first IP from X-Forwarded-For
					if ips := strings.Split(xForwardedFor, ","); len(ips) > 0 {
						clientIP = strings.TrimSpace(ips[0])
					}
				} else {
					// Fall back to RemoteAddr (remove port if present)
					clientIP = strings.Split(r.RemoteAddr, ":")[0]
				}

				allowIPs, isCheck := getIPs(ctx)
				if isCheck && !slices.Contains(allowIPs, clientIP) {
					return nil, fmt.Errorf("IP is not allow to access (%s)", clientIP)
				}
			}
			return handler(ctx, req)
		}
	}
}

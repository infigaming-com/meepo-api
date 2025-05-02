package middleware

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	mctx "github.com/infigaming-com/meepo-api/pkg/context"
)

func RequestInfoMiddleware() middleware.Middleware {
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

				requestInfo := mctx.RequestInfo{
					Method:     r.Method,
					Path:       r.URL.Path,
					Host:       r.Host,
					RemoteAddr: r.RemoteAddr,
					Country:    r.Header.Get("Cf-Ipcountry"),
					Origin:     r.Header.Get("Origin"),
					Referer:    r.Header.Get("Referer"),
					UserAgent:  r.Header.Get("User-Agent"),
					ClientIP:   clientIP,
				}
				ctx = mctx.WithRequestInfo(ctx, requestInfo)
			}
			return handler(ctx, req)
		}
	}
}

package middleware

import (
	"context"
	"io"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func LoggingMiddleware(enabled bool, lg log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				reply, err = handler(ctx, req)

				if enabled {
					log.NewHelper(lg).Debugw(
						"method", r.Method,
						"path", r.URL.Path,
						"host", r.Host,
						"remote_addr", r.RemoteAddr,
						"headers", r.Header,
						"query", r.URL.Query(),
						"request_body", readBody(r.Body),
						"status", getStatusCode(err),
						"response_body", reply,
						"error", err,
					)
				}
			}

			return reply, err
		}
	}
}

// responseWriter wraps http.ResponseWriter to capture response
type responseWriter struct {
	khttp.ResponseWriter
	status int
	body   []byte
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// readBody reads and restores the request body
func readBody(r io.ReadCloser) string {
	if r == nil {
		return ""
	}
	body, err := io.ReadAll(r)
	if err != nil {
		return ""
	}
	r.Close()
	return string(body)
}

// getStatusCode returns the HTTP status code based on the error
func getStatusCode(err error) int {
	if err == nil {
		return 200
	}
	if e, ok := err.(*errors.Error); ok {
		return int(e.Code)
	}
	return 500
}

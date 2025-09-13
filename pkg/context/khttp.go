package context

import (
	"context"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func IsRestful(ctx context.Context) bool {
	_, ok := khttp.RequestFromServerContext(ctx)
	return ok
}

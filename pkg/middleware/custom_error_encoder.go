package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func CustomErrorEncoder(errorReasonMap map[string]int32) khttp.EncodeErrorFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		if err == nil {
			return
		}

		se := errors.FromError(err)

		code, ok := errorReasonMap[se.Reason]
		if !ok {
			code = 500
		}

		resp := map[string]any{
			"code":    code,
			"message": se.Message,
			"reason":  se.Reason,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(resp)
	}
}

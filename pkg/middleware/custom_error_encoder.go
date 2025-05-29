package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
)

func CustomErrorEncoder(w http.ResponseWriter, r *http.Request, err error, errorReasonMap map[string]int32) {
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
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

package common

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/khakim88/test-promo/common/constant"
)

type successResponse struct {
	Success bool `json:"success"`
}

type dataResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	var (
		code    int
		message string
		data    map[string]interface{}
	)

	switch parsedErr := err.(type) {
	case *Error:
		code = parsedErr.CodeErr()
		message = parsedErr.Error()
		data = parsedErr.DataErr()
	default:
		code = ErrInternal.CodeErr()
		message = ErrInternal.Error()
	}

	statusCode := code / 100

	w.Header().Set(constant.CONTENTTYPEKEY, constant.CONTENTTYPEJSONANDCHARSETUTF8VALUE)
	if statusCode == http.StatusUnauthorized {
		w.Header().Set("Www-Authenticate", `Basic realm="Authorization Required"`)
	}
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(errorResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// EncodeLegacyError is used to mirror MPG1's error response format
func EncodeLegacyError(_ context.Context, err error, w http.ResponseWriter) {
	parsedErr, ok := err.(*Error)
	if ok {
		if parsedErr.Code/100 == http.StatusUnauthorized {
			w.Header().Set("Connection", `close`)
			w.Header().Set("X-Content-Type-Options", `nosniff`)
			w.Header().Set("Www-Authenticate", `Basic realm="Authorization Required"`)
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}
	}

	w.Header().Set(constant.CONTENTTYPEKEY, constant.CONTENTTYPEJSONANDCHARSETUTF8VALUE)

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func EncodeSuccessResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(constant.CONTENTTYPEKEY, constant.CONTENTTYPEJSONANDCHARSETUTF8VALUE)

	resp := successResponse{Success: true}
	return json.NewEncoder(w).Encode(resp)
}

func EncodeResponseWithData(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(constant.CONTENTTYPEKEY, constant.CONTENTTYPEJSONANDCHARSETUTF8VALUE)

	resp := dataResponse{Data: response}
	return json.NewEncoder(w).Encode(resp)
}

func EncodeResponseNoData(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func EncodeResponseWithCount(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(constant.CONTENTTYPEKEY, constant.CONTENTTYPEJSONANDCHARSETUTF8VALUE)
	return json.NewEncoder(w).Encode(response)
}

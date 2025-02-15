package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mvndaai/ctxerr"
	ctxerrhttp "github.com/mvndaai/ctxerr/http"
	"github.com/mvndaai/known-anywhere/internal/config"
)

type (
	Return struct {
		Success bool `json:"success"`

		// These should all be json objects
		Data  any                 `json:"data,omitempty"`
		Error *ctxerrhttp.Details `json:"error,omitempty"`
		Meta  any                 `json:"meta,omitempty"`
	}
)

type GenericHandlerFunc func(r *http.Request) (data, meta any, status int, _ error)

func GenericToHTTP(handler GenericHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ret Return
		data, meta, status, err := handler(r)
		if err != nil {
			ctxerr.Handle(err)
			debugErrors := config.Get().DebugErrors
			var errorResp ctxerrhttp.ErrorResponse
			status, errorResp = ctxerrhttp.StatusCodeAndResponse(err, debugErrors, debugErrors)
			ret.Error = &errorResp.Error
		} else {
			ret.Success = true
			ret.Data = data
			ret.Meta = meta
		}

		encoder := json.NewEncoder(w)
		indent, _ := strconv.ParseBool(r.Header.Get("Indent"))
		if indent {
			encoder.SetIndent("", "\t")
		}
		err = encoder.Encode(ret)
		if err != nil {
			ctxerr.Handle(ctxerr.Wrap(r.Context(), err, "8e9ba72c-7279-42bd-b01d-7d453b7915a3", "writing response"))
		}

		// TODO figure out why this is always 200
		if status != 0 && status != http.StatusOK {
			w.WriteHeader(status)
		}
	}
}

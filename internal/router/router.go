package router

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/mvndaai/ctxerr"
	ctxerrhttp "github.com/mvndaai/ctxerr/http"
)

var routes = map[string][]string{}

type (
	Return struct {
		Success bool `json:"success"`

		// These should all be json objects
		Data  any                 `json:"data,omitempty"`
		Error *ctxerrhttp.Details `json:"error,omitempty"`
		Meta  any                 `json:"meta,omitempty"`
	}
)

type WraperFunc (func(r *http.Request) (data, meta any, status int, _ error))

func NewRoute(method, path string, wf WraperFunc, middleware ...func(http.Handler) http.Handler) {
	if _, ok := routes[method]; !ok {
		routes[method] = make([]string, 0)
	}
	routes[method] = append(routes[method], path)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO add options
		if r.Method != method {
			return
		}
		var ret Return
		data, meta, status, err := wf(r)
		if err != nil {
			ctxerr.Handle(err)
			debugErrors, _ := strconv.ParseBool(os.Getenv("DEBUG_ERRORS"))
			var errorResp ctxerrhttp.ErrorResponse
			status, errorResp = ctxerrhttp.StatusCodeAndResponse(err, debugErrors, debugErrors)
			ret.Error = &errorResp.Error
		} else {
			ret.Success = true
			ret.Data = data
			ret.Meta = meta
		}

		err = json.NewEncoder(w).Encode(ret)
		if err != nil {
			ctxerr.Handle(ctxerr.Wrap(r.Context(), err, "8e9ba72c-7279-42bd-b01d-7d453b7915a3", "writing response"))
		}

		// TODO figure out why this is always 200
		if status != 0 && status != http.StatusOK {
			w.WriteHeader(status)
		}
	}))

	for _, m := range middleware {
		handler = m(handler)
	}

	http.Handle(path, handler)
}

func ListRoutesHandler(r *http.Request) (data, meta any, status int, _ error) {
	return routes, nil, http.StatusOK, nil
}

func GetPort() string {
	if v := os.Getenv("PORT"); v != "" {
		return ":" + v
	}
	return ":8080"
}

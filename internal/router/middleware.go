package router

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/mvndaai/ctxerr"
	ctxerrhttp "github.com/mvndaai/ctxerr/http"
	"github.com/mvndaai/known-socially/internal/jwt"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := func() error {
			claims, err := jwt.GetJWTClaims(r)
			if err != nil {
				return ctxerr.WrapHTTP(ctx, err, "2a321992-5c73-4c2f-b55b-6c291984e1f7", "invalid auth token", http.StatusUnauthorized, "invalid token format")
			}
			err = claims.EnsureClaims(ctx, r.Method, r.URL.Path, r.URL.Query(), "")
			if err != nil {
				return ctxerr.WrapHTTP(ctx, err, "d0257175-2fbd-46e8-9ef5-9dc1212c5491", "failed jwt claims", http.StatusUnauthorized, "failed to ensure claims")
			}
			return nil
		}()
		if err != nil {
			ctxerr.Handle(err)
			debugErrors, _ := strconv.ParseBool(os.Getenv("DEBUG_ERRORS"))
			status, errorResp := ctxerrhttp.StatusCodeAndResponse(err, debugErrors, debugErrors)
			b, _ := json.Marshal(errorResp)
			http.Error(w, string(b), status)
			return
		}

		next.ServeHTTP(w, r)
	})
}

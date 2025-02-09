package router

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
	ctxerrhttp "github.com/mvndaai/ctxerr/http"
	"github.com/mvndaai/known-socially/internal/config"
	"github.com/mvndaai/known-socially/internal/db"
	"github.com/mvndaai/known-socially/internal/jwt"
	"github.com/mvndaai/known-socially/internal/types"
)

func JWTMiddleware(pg *db.DB) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
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
				if pg != nil {
					jwtID, err := uuid.Parse(claims.ID)
					if err != nil {
						return ctxerr.WrapHTTP(ctx, err, "a938b89e-5c87-4078-a342-68e6ab643d59", "JWT ID not a uuid", http.StatusBadRequest, "jwt id not uuid")
					}
					users, _, err := pg.ListLogout(ctx, types.Logout{JWTID: jwtID}, types.Pagination{})
					if err != nil {
						return ctxerr.QuickWrap(ctx, err)
					}
					if len(users) > 0 {
						return ctxerr.NewHTTP(ctx, "bedd9e82-afb5-4593-989c-bd2a02b03265", "jwt logged out", http.StatusUnauthorized, "jwt id logged out")
					}
				}
				return nil
			}()
			if err != nil {
				ctxerr.Handle(err)
				debugErrors := config.DebugErrors()
				status, errorResp := ctxerrhttp.StatusCodeAndResponse(err, debugErrors, debugErrors)
				b, _ := json.Marshal(errorResp)
				http.Error(w, string(b), status)
				return
			}
			next.ServeHTTP(w, r)
		}
	}
}

func CleanUpParamsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Clean up the URL parameters
		nq := url.Values{}
		for k, vs := range r.URL.Query() {
			var nvs []string
			for _, v := range vs {
				nvs = append(nvs, strings.TrimSpace(v))
			}
			nq[strings.TrimSpace(strings.ToLower(k))] = nvs
		}
		r.URL.RawQuery = nq.Encode()
		next.ServeHTTP(w, r.WithContext(r.Context()))
	}
}

func JWTSubjectMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := jwt.GetJWTClaims(r)
		if err == nil && claims != nil {
			ctx := jwt.ContextWithSubject(r.Context(), claims.Subject)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	}
}

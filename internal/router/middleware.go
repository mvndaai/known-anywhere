package router

import (
	"net/http"

	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/jwt"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, err := jwt.GetJWTClaims(r)
		if err != nil {
			err := ctxerr.New(ctx, "2a321992-5c73-4c2f-b55b-6c291984e1f7", "invalid token format")
			ctxerr.Handle(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		err = claims.EnsureClaims(ctx, r.Method, r.URL.Path, r.URL.Query(), "")
		if err != nil {
			err = ctxerr.Wrap(ctx, err, "d0257175-2fbd-46e8-9ef5-9dc1212c5491", "failed to ensure claims")
			ctxerr.Handle(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

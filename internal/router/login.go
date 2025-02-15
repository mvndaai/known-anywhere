package router

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-anywhere/internal/jwt"
	"github.com/mvndaai/known-anywhere/internal/types"
)

func loginsHandler(r *http.Request) (data, meta any, status int, _ error) {
	return nil, nil, http.StatusBadGateway, ctxerr.New(r.Context(), "82bd850e-23aa-4fe0-8b9b-801aab34de0a", "test error")
}

func (h *Handler) logoutHandler(r *http.Request) (data, meta any, status int, _ error) {
	ctx := r.Context()
	claims, err := jwt.GetJWTClaims(r)
	if err != nil {
		return nil, nil, 0, ctxerr.QuickWrap(ctx, err)
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, nil, 0, ctxerr.WrapHTTP(ctx, err, "f57df46c-d870-4ce1-a2c2-86feefe792ae", "JWT ID not a uuid", http.StatusBadRequest, "jwt id not uuid")
	}
	et, err := claims.GetExpirationTime()
	if err != nil {
		return nil, nil, 0, ctxerr.WrapHTTP(ctx, err, "a2f2d73d-e1de-4b36-82d7-2bf78abd5033", "JWT expiration time not a valid time", http.StatusBadRequest, "jwt expiration time not valid time")
	}
	_, err = h.db.LogoutCreate(ctx, types.Logout{
		JWTID:      id,
		Expiration: et.Time,
	})
	if err != nil {
		return nil, nil, 0, ctxerr.QuickWrap(ctx, err)
	}

	return nil, nil, http.StatusOK, nil
}

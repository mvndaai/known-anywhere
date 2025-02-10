package db

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/jwt"
	"github.com/mvndaai/known-socially/internal/types"
)

const (
	tableLogouts = "logouts"
)

func scanLogout(scanner interface {
	Scan(dest ...any) error
}) (types.Logout, error) {
	var d types.Logout

	var id uuid.UUID
	err := scanner.Scan(
		&id,
		&d.JWTID,
		&d.Expiration,
		&d.UserID,
	)
	_ = id // ignore id because the type don't have it
	return d, err
}

func (v *DB) LogoutCreate(ctx context.Context, l types.Logout) (uuid.UUID, error) {
	l.UserID = jwt.SubjectFromContext(ctx)
	v.cache.DeleteJWTLogout(l.UserID)
	id, err := insertAndReturnID(ctx, v.db, tableLogouts, l)
	return id, ctxerr.QuickWrap(ctx, err)
}

func (v *DB) ListLogout(ctx context.Context, filters types.Logout, pagination types.Pagination) ([]types.Logout, types.PaginationResponse, error) {
	vs, pg, err := listItems(ctx, v.db, tableLogouts, filters, pagination, nil,
		func(rows *sql.Rows) (types.Logout, error) {
			return scanLogout(rows)
		})
	return vs, pg, ctxerr.QuickWrap(ctx, err)
}

func (v *DB) JWTAllowed(ctx context.Context, userID, jwtID string) error {
	userIDUUID, err := uuid.Parse(userID)
	if err != nil {
		return ctxerr.WrapHTTP(ctx, err, "9db70054-dec1-4518-8fa9-37794220931d", "User ID not a uuid", http.StatusBadRequest, "user id not uuid")
	}
	logouts, ok := v.cache.GetJWTLogout(userIDUUID)
	if !ok {
		var err error
		logouts, _, err = listItems(ctx, v.db, tableLogouts, types.Logout{UserID: userIDUUID}, types.Pagination{Limit: 100}, []Wheres{{where: "expiration < now()"}},
			func(rows *sql.Rows) (types.Logout, error) {
				return scanLogout(rows)
			})
		if err != nil {
			return ctxerr.QuickWrap(ctx, err)
		}
		v.cache.SetJWTLogout(userIDUUID, logouts)
	}
	if len(logouts) == 0 {
		return nil
	}

	jwtIDUUID, err := uuid.Parse(jwtID)
	if err != nil {
		return ctxerr.WrapHTTP(ctx, err, "a938b89e-5c87-4078-a342-68e6ab643d59", "JWT ID not a uuid", http.StatusBadRequest, "jwt id not uuid")
	}
	for _, lg := range logouts {
		if lg.JWTID == jwtIDUUID {
			return ctxerr.NewHTTP(ctx, "e7d0cd66-29f6-45cb-9b4e-a8fbd4d329b8", "jwt logged out", http.StatusUnauthorized, "jwt id logged out")
		}
	}
	return nil
}

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/mvndaai/known-socially/internal/jwt"
	"github.com/mvndaai/known-socially/internal/types"
)

const (
	tableLogout = "logout"
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
	)
	_ = id // ignore id because the type don't have it
	return d, err
}

func (pg *Postgres) LogoutCreate(ctx context.Context, l types.Logout) (uuid.UUID, error) {
	userID := jwt.SubjectFromContext(ctx)
	return insertAndReturnID(ctx, pg.db, tableLogout, l, columnValue{name: "user_id", value: userID})
}

func (pg *Postgres) ListLogout(ctx context.Context, filters types.Logout, pagination types.Pagination) ([]types.Logout, types.PaginationResponse, error) {
	return listItems(ctx, pg.db, tableLogout, filters, pagination,
		func(rows *sql.Rows) (types.Logout, error) {
			return scanLogout(rows)
		})
}

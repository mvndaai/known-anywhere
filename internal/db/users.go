package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/mvndaai/known-socially/internal/types"
)

const (
	tableUsers = "users"
)

func scanUser(scanner interface {
	Scan(dest ...any) error
}) (types.User, error) {
	var u types.User
	err := scanner.Scan(
		&u.ID,
		NullableScan(func(v string) { u.Username = v }),
		NullableScan(func(v string) { u.DisplayName = v }),
	)
	return u, err
}

func (pg *Postgres) CreateUser(ctx context.Context, u types.UserCreate) (uuid.UUID, error) {
	return insertAndReturnID(ctx, pg.db, tableUsers, u)
}

func (pg *Postgres) GetUser(ctx context.Context, id string) (types.User, error) {
	return get(ctx, pg.db, tableUsers, id, scanUser)
}

func (pg *Postgres) ListUsers(ctx context.Context, filters types.UserCreate, pagination types.Pagination) ([]types.User, types.PaginationResponse, error) {
	return listItems(ctx, pg.db, tableUsers, filters, pagination,
		func(rows *sql.Rows) (types.User, error) {
			return scanUser(rows)
		})
}

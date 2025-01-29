package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/mvndaai/known-socially/internal/jwt"
	"github.com/mvndaai/known-socially/internal/types"
)

const (
	tableDomains = "domains"
)

func scanDomain(scanner interface {
	Scan(dest ...any) error
}) (types.Domain, error) {
	var d types.Domain
	err := scanner.Scan(
		&d.ID,
		NullableScan(func(v string) { d.DisplayName = v }),
		NullableScan(func(v string) { d.Description = v }),
		NullableScan(func(v string) { d.Notes = v }),
	)
	return d, err
}

func (pg *Postgres) CreateDomain(ctx context.Context, d types.DomainCreate) (uuid.UUID, error) {
	creator := jwt.SubjectFromContext(ctx)
	return insertAndReturnID(ctx, pg.db, tableDomains, d, columnValue{name: "creator", value: creator})
}

func (pg *Postgres) GetDomain(ctx context.Context, id string) (types.Domain, error) {
	return get(ctx, pg.db, tableDomains, id, scanDomain)
}

func (pg *Postgres) ListDomains(ctx context.Context, filters types.DomainCreate, pagination types.Pagination) ([]types.Domain, types.PaginationResponse, error) {
	return listItems(ctx, pg.db, tableDomains, filters, pagination,
		func(rows *sql.Rows) (types.Domain, error) {
			return scanDomain(rows)
		})
}

package db

import (
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/types"
)

func (pg *Postgres) CreateDomain(ctx context.Context, d types.DomainCreate) (uuid.UUID, error) {
	// TODO pass through the JWT user

	creator := uuid.UUID{}

	result, err := pg.db.ExecContext(ctx, `
		INSERT INTO domains (
			display_name,
			description,
			notes,
			creator
		) VALUES ($1,$2,$3,$4)
	`, d.DisplayName, d.Description, d.Notes, creator)
	if err != nil {
		ctx = ctxerr.SetField(ctx, "body", d)
		return uuid.UUID{}, ctxerr.Wrap(ctx, err, "c486d504-230e-4fa7-9aea-f267b07fac50")
	}
	id, err := result.LastInsertId()
	log.Println(id)

	return uuid.UUID{}, nil
}

func (pg *Postgres) GetDomain(ctx context.Context, id string) (types.Domain, error) {
	var d types.Domain
	row := pg.db.QueryRowContext(ctx, `
		SELECT
			id,
			display_name,
			description,
			notes
		FROM domains
		WHERE id = $1
	`, id)
	err := row.Scan(
		&d.ID,
		NullableScan(func(v string) { d.DisplayName = v }),
		NullableScan(func(v string) { d.Description = v }),
		NullableScan(func(v string) { d.Notes = v }),
	)
	if err != nil {
		return d, ctxerr.Wrap(ctx, err, "7e644c35-7289-494c-8ee3-856edcc0b5bd")
	}
	return d, nil
}

func (pg *Postgres) ListDomains(ctx context.Context, l types.DomainList) ([]types.Domain, types.PaginationResponse, error) {
	pr := types.PaginationResponse{}

	wheres := []string{}
	args := []any{}
	vc := varCount{}
	if l.Filters.DisplayName != "" {
		wheres = append(wheres, "display_name = "+vc.Next())
		args = append(args, l.Filters.DisplayName)
	}
	if l.Filters.Description != "" {
		wheres = append(wheres, "description = "+vc.Next())
		args = append(args, l.Filters.Description)
	}
	if l.Filters.Notes != "" {
		wheres = append(wheres, "notes = "+vc.Next())
		args = append(args, l.Filters.Notes)
	}

	if l.Pagination.Cursor != "" {
		wheres = append(wheres, "id > "+vc.Next())
		args = append(args, l.Pagination.Cursor)
	}
	where := strings.Join(wheres, " AND ")
	if where != "" {
		where = "WHERE " + where
	}

	err := pg.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM domains `+where, args...).Scan(&pr.Total)
	if err != nil {
		return nil, pr, ctxerr.Wrap(ctx, err, "c5c072fe-8e87-47be-9e15-ec390dfc8d35")
	}

	where += " ORDER BY id ASC LIMIT " + vc.Next()
	args = append(args, l.Pagination.Limit)

	query := `
		SELECT
			id,
			display_name,
			description,
			notes
		FROM domains
	` + where
	log.Println(query, args)

	rows, err := pg.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, pr, ctxerr.Wrap(ctx, err, "4f0f387a-3c09-4170-87f8-ed6976e3cfcd")
	}
	defer rows.Close()

	var domains []types.Domain
	for rows.Next() {
		var d types.Domain
		err = rows.Scan(
			&d.ID,
			NullableScan(func(v string) { d.DisplayName = v }),
			NullableScan(func(v string) { d.Description = v }),
			NullableScan(func(v string) { d.Notes = v }),
		)
		if err != nil {
			return nil, pr, ctxerr.Wrap(ctx, err, "f5a3c9c5-7f9d-4a3f-9d4c-4c9d3b6c0c8d")
		}
		domains = append(domains, d)
	}

	pr.Cursor = domains[len(domains)-1].ID.String()
	return domains, pr, nil
}

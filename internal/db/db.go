package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/types"
)

type Postgres struct {
	db *sql.DB
}

func (pg *Postgres) Connect(ctx context.Context) error {
	// Connect to the database

	un := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("POSTGRES_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}
	// TODO add POSTGRES_HOST and POSTGRES_PORT

	dbSorceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", un, pw, dbname, sslmode)
	var err error
	pg.db, err = sql.Open("postgres", dbSorceName)
	if err != nil {
		return ctxerr.Wrap(ctx, err, "e94bf5b7-5449-41a6-ae92-2103fa475845", "Failed to connect to the database")
	}

	err = pg.CreateTables(ctx)
	if err != nil {
		return ctxerr.QuickWrap(ctx, err)
	}

	return nil
}

func (pg *Postgres) Close(ctx context.Context) error {
	// Close the database connection
	err := pg.db.Close()
	if err != nil {
		return ctxerr.Wrap(ctx, err, "5f050fa7-37ec-4ba6-8268-aa8ae03dc045", "Failed to close the database connection")
	}
	return nil
}

type varCount struct {
	i int
}

func (vc *varCount) Next() string {
	vc.i++
	return fmt.Sprintf("$%d", vc.i)
}

type nullable[T any] struct {
	value T
	fn    func(T)
}

func (n *nullable[T]) Scan(value interface{}) error {
	if value == nil {
		var zero T
		n.fn(zero)
		return nil
	}
	switch v := value.(type) {
	case T:
		n.fn(v)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

func NullableScan[T any](fn func(T)) *nullable[T] {
	return &nullable[T]{fn: fn}
}

func getSelectFields[T any]() string {
	var t T
	typ := reflect.TypeOf(t)
	fields := []string{"id"}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous {
			// For embedded structs, get their fields
			embedded := reflect.New(field.Type).Elem().Interface()
			embType := reflect.TypeOf(embedded)
			for j := 0; j < embType.NumField(); j++ {
				if tag := embType.Field(j).Tag.Get("json"); tag != "" {
					fields = append(fields, tag)
				}
			}
		} else if tag := field.Tag.Get("json"); tag != "" && tag != "id" {
			fields = append(fields, tag)
		}
	}
	return strings.Join(fields, ", ")
}

func listItems[T any, F any](
	ctx context.Context,
	db *sql.DB,
	tableName string,
	filters F,
	pagination types.Pagination,
	scan func(*sql.Rows) (T, error),
) ([]T, types.PaginationResponse, error) {
	pr := types.PaginationResponse{}
	vc := varCount{}

	selectFields := getSelectFields[T]()

	// Build where clause using reflection
	wheres := []string{}
	args := []any{}
	v := reflect.ValueOf(filters)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Skip empty values
		if value.IsZero() {
			continue
		}

		columnName := types.JSONTag(filters, field.Name)
		if columnName == "" {
			continue
		}

		wheres = append(wheres, columnName+" = "+vc.Next())
		args = append(args, value.Interface())
	}

	if pagination.Cursor != "" {
		wheres = append(wheres, "id > "+vc.Next())
		args = append(args, pagination.Cursor)
	}
	where := strings.Join(wheres, " AND ")
	if where != "" {
		where = "WHERE " + where
	}

	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM "+tableName+" "+where, args...).Scan(&pr.Total)
	if err != nil {
		return nil, pr, ctxerr.Wrap(ctx, err, "c5c072fe-8e87-47be-9e15-ec390dfc8d35")
	}

	where += " ORDER BY id ASC LIMIT " + vc.Next()
	args = append(args, pagination.Limit)

	query := fmt.Sprintf(`SELECT %s FROM %s %s`, selectFields, tableName, where)
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, pr, ctxerr.Wrap(ctx, err, "4f0f387a-3c09-4170-87f8-ed6976e3cfcd")
	}
	defer rows.Close()

	var items []T
	for rows.Next() {
		item, err := scan(rows)
		if err != nil {
			return nil, pr, err
		}
		items = append(items, item)
	}

	if len(items) == pagination.Limit {
		if last, ok := any(items[len(items)-1]).(interface{ GetID() uuid.UUID }); ok {
			pr.Cursor = last.GetID().String()
		}
	}
	return items, pr, nil
}

func getInsertColumns(item any, additionalColumns ...string) (columns []string, args []any) {
	typ := reflect.TypeOf(item)
	val := reflect.ValueOf(item)

	for i := 0; i < typ.NumField(); i++ {
		if tag := typ.Field(i).Tag.Get("json"); tag != "" && tag != "id" {
			columns = append(columns, tag)
			args = append(args, val.Field(i).Interface())
		}
	}

	for _, col := range additionalColumns {
		columns = append(columns, col)
	}
	return columns, args
}

func getDollarSigns(count int) []string {
	signs := make([]string, count)
	for i := range signs {
		signs[i] = fmt.Sprintf("$%d", i+1)
	}
	return signs
}

type columnValue struct {
	name  string
	value any
}

func insertAndReturnID[T any](ctx context.Context, db *sql.DB, tableName string, item T, additionalCols ...columnValue) (uuid.UUID, error) {
	columns, args := getInsertColumns(item)

	// Add additional columns and their values
	for _, col := range additionalCols {
		columns = append(columns, col.name)
		args = append(args, col.value)
	}

	dollars := getDollarSigns(len(columns))

	var id uuid.UUID
	err := db.QueryRowContext(ctx, fmt.Sprintf(`
		INSERT INTO %s (
			%s
		) VALUES (%s)
		RETURNING id
	`, tableName, strings.Join(columns, ",\n\t\t\t"), strings.Join(dollars, ",")),
		args...).Scan(&id)

	if err != nil {
		ctx = ctxerr.SetField(ctx, "body", item)
		return uuid.UUID{}, ctxerr.Wrap(ctx, err, "c486d504-230e-4fa7-9aea-f267b07fac50")
	}

	return id, nil
}

func get[T any](
	ctx context.Context,
	db *sql.DB,
	tableName string,
	id string,
	scan func(scanner interface{ Scan(dest ...any) error }) (T, error),
) (T, error) {
	fields := getSelectFields[T]()
	item, err := scan(db.QueryRowContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM %s
		WHERE id = $1
	`, fields, tableName), id))
	if err != nil {
		return item, ctxerr.Wrap(ctx, err, "7e644c35-7289-494c-8ee3-856edcc0b5bd")
	}
	return item, nil
}

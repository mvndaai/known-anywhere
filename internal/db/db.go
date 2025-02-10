package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/config"
	"github.com/mvndaai/known-socially/internal/types"
)

type DB struct {
	db    *sql.DB
	cache *CacheImpl
}

func New(ctx context.Context) (*DB, error) {
	postgress, err := sql.Open("postgres", config.Get().Postgres.DataSourceName())
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "e94bf5b7-5449-41a6-ae92-2103fa475845", "Failed to connect to postgres")
	}
	ret := &DB{
		db:    postgress,
		cache: newCache(),
	}

	// Create tables if they don't exist
	if bl, _ := strconv.ParseBool(os.Getenv("CREATE_TABLES")); bl {
		err = ret.CreateTables(ctx)
		if err != nil {
			return nil, ctxerr.QuickWrap(ctx, err)
		}
	}

	return ret, nil
}

func (v *DB) Close(ctx context.Context) error {
	// Close the database connection
	err := v.db.Close()
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

type whereClause struct {
	varCount int
	wheres   []string
	args     []any
}

func (wc *whereClause) Add(where string, arg any) {
	if arg == nil {
		wc.wheres = append(wc.wheres, where)
		return
	}
	wc.wheres = append(wc.wheres, where+wc.nextVar())
	wc.args = append(wc.args, arg)
}

func (wc *whereClause) WhereAndArgs() (string, []any) {
	if len(wc.wheres) == 0 {
		return "", nil
	}
	return "WHERE " + strings.Join(wc.wheres, " AND "), wc.args
}

func (wc *whereClause) nextVar() string {
	wc.varCount++
	return fmt.Sprintf("$%d", wc.varCount)
}

type Wheres struct {
	where string
	arg   any
}

func listItems[T any, F any](
	ctx context.Context,
	db *sql.DB,
	tableName string,
	filters F,
	pagination types.Pagination,
	wheres []Wheres,
	scan func(*sql.Rows) (T, error),
) ([]T, types.PaginationResponse, error) {
	pagination.Normalize()
	pr := types.PaginationResponse{}
	wc := whereClause{}
	for _, w := range wheres {
		wc.Add(w.where, w.arg)
	}

	selectFields := getSelectFields[T]()

	// Build where clause using reflection
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

		wc.Add(fmt.Sprintf("%s.%s = ", tableName, columnName), value.Interface())
	}

	where, args := wc.WhereAndArgs()
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM "+tableName+" "+where, args...).Scan(&pr.Total)
	if err != nil {
		return nil, pr, ctxerr.Wrap(ctx, err, "c5c072fe-8e87-47be-9e15-ec390dfc8d35")
	}

	if pagination.Cursor != "" {
		wc.Add("id > ", pagination.Cursor)
		where, args = wc.WhereAndArgs()
	}
	where += " ORDER BY id ASC LIMIT " + wc.nextVar()
	args = append(args, pagination.Limit)
	query := fmt.Sprintf(`SELECT %s FROM %s %s`, selectFields, tableName, where)
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		ctx = ctxerr.SetField(ctx, "query", query)
		return nil, pr, ctxerr.Wrap(ctx, err, "1d3f4034-4dd1-4772-9db0-d56365f67f11")
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
	if items == nil { // Don't return null, return an empty slice
		items = []T{}
	}
	if l := len(items); l > 0 && l == pagination.Limit {
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
	query := fmt.Sprintf(`
		INSERT INTO %s (
			%s
		) VALUES (%s)
		RETURNING id
	`, tableName, strings.Join(columns, ",\n\t\t\t"), strings.Join(dollars, ","))

	err := db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		ctx = ctxerr.SetField(ctx, "body", item)
		ctx = ctxerr.SetField(ctx, "query", query)
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

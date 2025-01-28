package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/mvndaai/ctxerr"
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

	dbSorceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", un, pw, dbname, sslmode)
	var err error
	pg.db, err = sql.Open("postgres", dbSorceName)
	if err != nil {
		return ctxerr.Wrap(ctx, err, "e94bf5b7-5449-41a6-ae92-2103fa475845", "Failed to connect to the database")
	}

	//err = pg.CreateTables(ctx)
	//if err != nil {
	//	return ctxerr.QuickWrap(ctx, err)
	//}

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

package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"
	"time"

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
	//dbport := os.Getenv("POSTGRES_PORT")
	//if dbport == "" {
	//	dbport = "5432"
	//}
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

type (
	NullScannerValuer interface {
		driver.Valuer
		sql.Scanner
	}
	NullBase struct {
		NullScanner NullScannerValuer // The temporary dest &sql.Null*{} type, must be pointer
		Dest        interface{}       // The final dest, must be pointer
	}
)

func Nullable(d any) *NullBase {
	switch d.(type) {
	case *string:
		return &NullBase{Dest: d, NullScanner: &sql.NullString{}}
	case *int32:
		return &NullBase{Dest: d, NullScanner: &sql.NullInt32{}}
	case *int64, *int:
		return &NullBase{Dest: d, NullScanner: &sql.NullInt64{}}
	case *int16:
		return &NullBase{Dest: d, NullScanner: &sql.NullInt16{}}
	case *byte:
		return &NullBase{Dest: d, NullScanner: &sql.NullByte{}}
	case *float64:
		return &NullBase{Dest: d, NullScanner: &sql.NullFloat64{}}
	case *bool:
		return &NullBase{Dest: d, NullScanner: &sql.NullBool{}}
	case *time.Time:
		return &NullBase{Dest: d, NullScanner: &sql.NullTime{}}
	}
	panic("unknown type")
	//return &NullBase{Dest: d, NullScanner: &sql.Null{}} // TODO maybe figure out using this generic type
}

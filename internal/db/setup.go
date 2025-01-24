package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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

func (pg *Postgres) CreateTables(ctx context.Context) error {
	// https://postgresql.verite.pro/blog/2024/07/15/uuid-v7-pure-sql.html
	_, err := pg.db.ExecContext(ctx, `
		CREATE FUNCTION uuidv7() RETURNS uuid
		AS $$
		select encode(
			set_bit(
			set_bit(
				overlay(uuid_send(gen_random_uuid()) placing
			substring(int8send((extract(epoch from clock_timestamp())*1000)::bigint) from 3)
			from 1 for 6),
			52, 1),
			53, 1), 'hex')::uuid;
		$$ LANGUAGE sql volatile;
	`)
	if err != nil {
		err = ctxerr.Wrap(ctx, err, "b9e3f5e6-6e2c-4e5c-8c0b-1f4c4e8e5a2b", "failed to create pg function uuidv7")
		ctxerr.Handle(err) // this always errors after first creation
	}

	_, err = pg.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id uuid DEFAULT uuidv7(),
    		email VARCHAR NOT NULL,
    		PRIMARY KEY (id)
		)`)
	if err != nil {
		return ctxerr.Wrap(ctx, err, "930187a3-bb61-4f6f-ae1f-3d54b700aff0", "Failed to create tables")
	}
	return nil
}

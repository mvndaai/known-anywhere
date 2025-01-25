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
		CREATE OR REPLACE FUNCTION uuidv7() RETURNS uuid
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

	// Create trigger function to update modified column
	_, err = pg.db.ExecContext(ctx, `
		CREATE OR REPLACE FUNCTION update_modified_column()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.modified = CURRENT_TIMESTAMP;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`)
	if err != nil {
		return ctxerr.Wrap(ctx, err, "377da5ce-43ff-415e-a8a4-362e7c5350b7", "failed to create trigger function")
	}

	tables := map[string]string{
		"users": `(
			id uuid DEFAULT uuidv7() PRIMARY KEY,
			username VARCHAR NOT NULL UNIQUE,
			display_name VARCHAR,

			deleted BOOLEAN default false,
			created TIMESTAMP default CURRENT_TIMESTAMP,
			modified TIMESTAMP default CURRENT_TIMESTAMP
		)`,
		"delete_audit": `(
			id uuid DEFAULT uuidv7() PRIMARY KEY,
			table_name VARCHAR NOT NULL,
			row_id uuid NOT NULL,

			creator uuid NOT NULL references users(id),
			deleted TIMESTAMP default CURRENT_TIMESTAMP,
			created TIMESTAMP default CURRENT_TIMESTAMP,
			modified TIMESTAMP default CURRENT_TIMESTAMP,
			UNIQUE (table_name, row_id)
		)`,
		"groups": `(
			id uuid DEFAULT uuidv7() PRIMARY KEY,
			display_name VARCHAR,

			deleted BOOLEAN default false,
			creator uuid NOT NULL references users(id),
			created TIMESTAMP default CURRENT_TIMESTAMP,
			modified TIMESTAMP default CURRENT_TIMESTAMP
		)`,
		"domains": `(
			id uuid DEFAULT uuidv7() PRIMARY KEY,
			domain VARCHAR NOT NULL UNIQUE,
			display_name VARCHAR,

			deleted BOOLEAN default false,
			creator uuid NOT NULL references users(id),
			created TIMESTAMP default CURRENT_TIMESTAMP,
			modified TIMESTAMP default CURRENT_TIMESTAMP
		)`,
		"domain_links": `(
			id_a uuid NOT NULL references domains(id),
			id_b uuid NOT NULL references domains(id),

			deleted BOOLEAN default false,
			creator uuid NOT NULL references users(id),
			created TIMESTAMP default CURRENT_TIMESTAMP,
			modified TIMESTAMP default CURRENT_TIMESTAMP,
			PRIMARY KEY (id_a, id_b)
		)`,
		"socials": `(
			id uuid DEFAULT uuidv7() PRIMARY KEY,
			domain_id uuid NOT NULL references domains(id),
			username VARCHAR,
			user_id VARCHAR,
			group_id uuid NOT NULL references groups(id),

			deleted BOOLEAN default false,
			creator uuid NOT NULL references users(id),
			created TIMESTAMP default CURRENT_TIMESTAMP,
			modified TIMESTAMP default CURRENT_TIMESTAMP,
			constraint either_email check (username is not null or user_id is not null),
			UNIQUE (domain_id, username, user_id, group_id)
		)`,
		"social_votes": `(
			id uuid DEFAULT uuidv7() PRIMARY KEY,
			social_id uuid NOT NULL references socials(id),
			downvote BOOLEAN,

			deleted BOOLEAN default false,
			user_id uuid NOT NULL references users(id),
			created TIMESTAMP default CURRENT_TIMESTAMP,
			modified TIMESTAMP default CURRENT_TIMESTAMP,
			UNIQUE (social_id, user_id)
		);
		COMMENT ON COLUMN social_votes.downvote IS 'if false then upvote';`,
	}

	for name, table := range tables {
		q := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s %s", name, table)
		_, err := pg.db.ExecContext(ctx, q)
		if err != nil {
			ctx = ctxerr.SetField(ctx, "query", q)
			return ctxerr.Wrap(ctx, err, "930187a3-bb61-4f6f-ae1f-3d54b700aff0", "failed to create table:", name)
		}

		// Create trigger for users table
		_, err = pg.db.ExecContext(ctx, fmt.Sprintf(`
			CREATE OR REPLACE TRIGGER update_%s_modified
			BEFORE UPDATE ON %s
			FOR EACH ROW
			EXECUTE FUNCTION update_modified_column();
		`, name, name))
		if err != nil {
			return ctxerr.Wrap(ctx, err, "a2261f40-31bb-4a3b-abcf-c80f23f3c93a", "failed to create trigger for table", name)
		}
	}

	return nil
}

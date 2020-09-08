// Package postgres is used to perform operations against postgres databases
package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/stdlib"
	"github.com/luna-duclos/instrumentedsql"
	"github.com/luna-duclos/instrumentedsql/opentracing"

	"pkg.dsb.dev/health"
	"pkg.dsb.dev/metrics"

	// Migration driver for postgres.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

// Open opens a connection to an SQL database using the provided DSN. Database migrations are also performed.
func Open(dsn string) (*sql.DB, error) {
	dsn = os.ExpandEnv(dsn)
	drv := instrumentedsql.WrapDriver(
		stdlib.GetDefaultDriver(),
		instrumentedsql.WithTracer(opentracing.NewTracer(false)),
		instrumentedsql.WithOmitArgs(),
		instrumentedsql.WithOpsExcluded(instrumentedsql.OpSQLRowsNext),
	)

	conn, err := drv.OpenConnector(dsn)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(conn)

	health.AddCheck("database", db.Ping)
	metrics.AddSQLStats(db)
	return db, db.Ping()
}

// WithinTransaction invokes 'cb' within an SQL transaction. If the callback returns an error, the transaction
// is rolled back.
func WithinTransaction(ctx context.Context, db *sql.DB, cb func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := cb(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

// CreateDatabaseWithUser creates a new user-password combination as the owner of a desired
// new database.
func CreateDatabaseWithUser(ctx context.Context, db *sql.DB, name, user, pass string) error {
	const (
		userQueryFmt = "CREATE USER %s WITH PASSWORD '%s'"
		dbQueryFmt   = "CREATE DATABASE %s WITH OWNER %s"
		permQueryFmt = "GRANT ALL PRIVILEGES ON DATABASE %s TO %s"
	)

	queries := []string{
		fmt.Sprintf(userQueryFmt, user, pass),
		fmt.Sprintf(dbQueryFmt, name, user),
		fmt.Sprintf(permQueryFmt, name, user),
	}

	return WithinTransaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		for _, q := range queries {
			if _, err := db.ExecContext(ctx, q); err != nil {
				return err
			}
		}

		return nil
	})
}

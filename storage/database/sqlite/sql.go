// Package sqlite is used to perform operations against sqlite databases
package sqlite

import (
	"database/sql"
	"os"
	"time"

	"github.com/luna-duclos/instrumentedsql"
	"github.com/luna-duclos/instrumentedsql/opentracing"
	"github.com/mattn/go-sqlite3"

	"pkg.dsb.dev/health"
	"pkg.dsb.dev/metrics"
	"pkg.dsb.dev/multierror"
	"pkg.dsb.dev/storage/database"
)

const (
	maxLifetime  = time.Minute * 30
	maxIdleConns = 20
	maxOpenConns = 200
)

// Open opens a connection to an SQL database using the provided DSN. Migrations are performed if the source
// is non-nil.
func Open(dsn string, migrations *database.MigrationSource) (*sql.DB, error) {
	dsn = os.ExpandEnv(dsn)
	drv := instrumentedsql.WrapDriver(
		&sqlite3.SQLiteDriver{},
		instrumentedsql.WithTracer(opentracing.NewTracer(false)),
		instrumentedsql.WithOmitArgs(),
		instrumentedsql.WithOpsExcluded(instrumentedsql.OpSQLRowsNext),
	)

	conn, err := drv.OpenConnector(dsn)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(conn)
	db.SetConnMaxLifetime(maxLifetime)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	if migrations != nil {
		err = database.MigrateUp(migrations, db)
		if err != nil {
			return nil, multierror.Append(err, db.Close())
		}
	}

	health.AddCheck("sqlite", db.Ping)
	metrics.AddSQLStats(db)
	return db, db.Ping()
}

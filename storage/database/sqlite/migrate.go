package sqlite

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"

	pkgdb "pkg.dsb.dev/storage/database"
)

// NewDirectoryMigration creates a migration from the provided directory path.
// The parameter must be a directory.
func NewDirectoryMigration(dir string) *pkgdb.MigrationSource {
	return pkgdb.NewDirectoryMigration(newSQLiteDBDriver, dir)
}

// NewBindataMigration creates a migration from the provided bindata asset.
func NewBindataMigration(b *bindata.AssetSource) *pkgdb.MigrationSource {
	return pkgdb.NewBindataMigration(newSQLiteDBDriver, b)
}

func newSQLiteDBDriver(db *sql.DB) (database.Driver, error) {
	return sqlite3.WithInstance(db, new(sqlite3.Config))
}

package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func MigrateDB(db *sqlx.DB, dbName, migrationPath string) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to setup postgres driver")
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, dbName, driver)
	if err != nil {
		return errors.Wrap(err, "failed to read migrations")
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "failed to migrate postgres")
	}
	return nil
}

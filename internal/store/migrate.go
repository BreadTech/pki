package store

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/BreadTech/pki/internal/migrations"
)

// Useful for migrate pkg logging
type logWrapper struct {
	*logrus.Logger
}

func newLog() migrate.Logger {
	return &logWrapper{logrus.New()}
}

func (l *logWrapper) Verbose() bool {
	return true
}

// Just to silence the linter when logger is not used
var _ = newLog()

// RunMigrations will run all database migrations in pkg/migrations. Please see the
// README.md in that folder to add more migrations.
func RunMigrations(db *gorm.DB) error {
	migrationTarget, err := sqlite3.WithInstance(db.DB(), &sqlite3.Config{})
	if err != nil {
		logrus.WithError(err).Error("migrate: failed to create database.Driver")
		return err
	}

	migrationSource, err := bindata.WithInstance(bindata.Resource(
		migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		}))
	if err != nil {
		logrus.WithError(err).Error("migrate: failed to fetch assets")
		return err
	}
	m, err := migrate.NewWithInstance("go-bindata", migrationSource, "sqlite3", migrationTarget)
	// m.Log = newLog()
	if err != nil {
		logrus.WithError(err).Error("migrate: failed to create migrate instance")
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.WithError(err).Error("migrate: failed to execute migration")
		return err
	}
	return nil
}

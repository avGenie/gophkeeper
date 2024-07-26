package migrator

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

const (
	migrationDB     = "postgres"
	migrationFolder = "migrations"
)

// Migrator Performs database migrations
type Migrator struct {
	db *sql.DB
}

// NewMigrator Creates migrator object
func NewMigrator(dsn, driverName string, migrationFs embed.FS) (*Migrator, error) {
	goose.SetBaseFS(migrationFs)

	if err := goose.SetDialect(migrationDB); err != nil {
		return nil, fmt.Errorf("couldn't set goose dialect: %w", err)
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't open sql connection: %w", err)
	}

	return &Migrator{
		db: db,
	}, nil
}

// NewMigrator Performs migration
func (m *Migrator) Up() error {
	if err := goose.Up(m.db, migrationFolder); err != nil {
		return fmt.Errorf("couldn't up goose migration: %w", err)
	}

	return nil
}

// NewMigrator Rolls back migration
func (m *Migrator) Down() error {
	if err := goose.Down(m.db, migrationFolder); err != nil {
		return fmt.Errorf("couldn't down goose migration: %w", err)
	}

	return nil
}

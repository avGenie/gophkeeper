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

type Migrator struct {
	db *sql.DB
}

func NewMigrator(dsn string, migrationFs embed.FS) (*Migrator, error) {
	goose.SetBaseFS(migrationFs)

	if err := goose.SetDialect(migrationDB); err != nil {
		return nil, fmt.Errorf("couldn't set goose dialect: %w", err)
	}

	db, err := sql.Open("pgx/v5", dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't open sql connection: %w", err)
	}

	return &Migrator{
		db: db,
	}, nil
}

func (m *Migrator) Up() error {
	if err := goose.Up(m.db, migrationFolder); err != nil {
		return fmt.Errorf("couldn't up goose migration: %w", err)
	}

	return nil
}

func (m *Migrator) Down() error {
	if err := goose.Down(m.db, migrationFolder); err != nil {
		return fmt.Errorf("couldn't down goose migration: %w", err)
	}

	return nil
}

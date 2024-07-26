package postgres

import (
	"context"
	"embed"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/config"
	"github.com/avGenie/gophkeeper/server/internal/usecase/migrator"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const driverName = "pgx/v5"

//go:embed migrations/*.sql
var migrationFs embed.FS

// Postgres Struct for postgres storage connection and usage
type Postgres struct {
	pool *pgxpool.Pool
}

// NewStorage Creates new postgres storage from config
func NewStorage(config config.StorageConfig) (*Postgres, error) {
	poolConfig, err := pgxpool.ParseConfig(config.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse postgres config: %w", err)
	}

	poolConfig.MaxConns = config.MaxPool

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("couldn't create postgres pool from config: %w", err)
	}

	err = migration(pool.Config().ConnConfig.ConnString())
	if err != nil {
		return nil, fmt.Errorf("couldn't process postgres migration: %w", err)
	}

	return &Postgres{
		pool: pool,
	}, nil
}

// Close Close postgres connection
func (p *Postgres) Close() {
	if p.pool != nil {
		p.pool.Close()
	}
}

func migration(dsn string) error {
	migrator, err := migrator.NewMigrator(dsn, driverName, migrationFs)
	if err != nil {
		return fmt.Errorf("couldn't create migration object")
	}

	err = migrator.Up()
	if err != nil {
		localErr := fmt.Errorf("couldn't up migration: %w", err)

		if err = migrator.Down(); err != nil {
			return fmt.Errorf("couldn't down migration after failed up: %w", err)
		}

		return localErr
	}

	return nil
}

package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/app/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	migrationDB     = "postgres"
	migrationFolder = "migrations"
)

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
	goose.SetBaseFS(migrationFs)

	if err := goose.SetDialect(migrationDB); err != nil {
		return fmt.Errorf("couldn't set goose dialect: %w", err)
	}

	db, err := sql.Open("pgx/v5", dsn)
	if err != nil {
		return fmt.Errorf("couldn't open sql connection: %w", err)
	}

	if err := goose.Up(db, migrationFolder); err != nil {
		return fmt.Errorf("couldn't up goose migration: %w", err)
	}

	return nil
}

package postgres

import (
	"context"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/app/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

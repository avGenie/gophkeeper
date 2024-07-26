package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (p *Postgres) CreateUser(ctx context.Context, user entity.User) error {
	query := `INSERT INTO users VALUES($1, $2, $3)`

	_, err := p.pool.Exec(ctx, query, user.ID, user.Login, user.Password)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return storage.ErrUserExists
		}

		return fmt.Errorf("couldn't execute create user postgres request: %w", err)
	}

	return nil
}

func (p *Postgres) GetUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := `SELECT id, password FROM users WHERE login=$1`

	row := p.pool.QueryRow(ctx, query, user.Login)
	if row == nil {
		return entity.User{}, fmt.Errorf("%s while authorize user", QueryRowMessage)
	}

	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, storage.ErrUserNotFound
		}

		return entity.User{}, fmt.Errorf("error while processing response row in postgres while getting user: %w", err)
	}

	return user, nil
}

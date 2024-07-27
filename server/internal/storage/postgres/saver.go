package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// SaveLoginPasswordData Saves login-password data to postgres storage
func (p *Postgres) SaveLoginPasswordData(ctx context.Context, data entity.LoginPassword, userID entity.UserID) error {
	query := `INSERT INTO login_password_data VALUES($1, $2, $3, $4, $5)`

	_, err := p.pool.Exec(ctx, query, userID, data.Name, data.Login, data.Password, data.Metadata.Data)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return storage.ErrDataExists
		}

		return fmt.Errorf("couldn't execute save login-password data postgres request: %w", err)
	}

	return nil
}

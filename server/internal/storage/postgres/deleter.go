package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/jackc/pgx/v5"
)

// DeleteLoginPasswordData Deletes login-password data from postgres storage for user
func (p *Postgres) DeleteLoginPasswordData(ctx context.Context, name entity.DataName, userID entity.UserID) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("couldn't create transaction while deleting login-password data postgres: %w", err)
	}
	defer tx.Rollback(ctx)

	querySelect := `SELECT name, login, password, meta FROM login_password_data WHERE user_login=$1 AND name=$2 FOR UPDATE`

	row := tx.QueryRow(ctx, querySelect, userID, name)
	if row == nil {
		return fmt.Errorf("%s while getting login-password data", queryRowMessage)
	}

	var outData entity.LoginPassword
	err = row.Scan(&outData.Name, &outData.Login, &outData.Password, &outData.Metadata.Data)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.ErrLoginPasswordDataNotFound
		}

		return fmt.Errorf("error while processing response row in postgres while getting login-password data: %w", err)
	}

	queryDelete := `DELETE FROM login_password_data WHERE user_login=$1 AND name=$2`

	_, err = tx.Exec(ctx, queryDelete, userID, name)
	if err != nil {
		return fmt.Errorf("couldn't delete login-password data from postgres request: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("couldn't commit transaction while deleting login-password data postgres: %w", err)
	}

	return nil
}

// DeleteTextData Deletes text data from postgres storage for user
func (p *Postgres) DeleteTextData(ctx context.Context, name entity.DataName, userID entity.UserID) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("couldn't create transaction while deleting text data postgres: %w", err)
	}
	defer tx.Rollback(ctx)

	querySelect := `SELECT name, text, meta FROM text_data WHERE user_login=$1 AND name=$2 FOR UPDATE`

	row := tx.QueryRow(ctx, querySelect, userID, name)
	if row == nil {
		return fmt.Errorf("%s while getting text data", queryRowMessage)
	}

	var outData entity.TextData
	err = row.Scan(&outData.Name, &outData.Text, &outData.Metadata.Data)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.ErrTextDataNotFound
		}

		return fmt.Errorf("error while processing response row in postgres while getting login-password data: %w", err)
	}

	queryDelete := `DELETE FROM text_data WHERE user_login=$1 AND name=$2`

	_, err = tx.Exec(ctx, queryDelete, userID, name)
	if err != nil {
		return fmt.Errorf("couldn't delete text data from postgres request: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("couldn't commit transaction while deleting text data postgres: %w", err)
	}

	return nil
}

// DeleteCardData Deletes card data from postgres storage for user
func (p *Postgres) DeleteCardData(ctx context.Context, name entity.DataName, userID entity.UserID) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("couldn't create transaction while deleting card data postgres: %w", err)
	}
	defer tx.Rollback(ctx)

	querySelect := `SELECT name, number, expiration_month, expiration_year, code, cardholder, meta FROM cards_data WHERE user_login=$1 AND name=$2 FOR UPDATE`

	row := tx.QueryRow(ctx, querySelect, userID, name)
	if row == nil {
		return fmt.Errorf("%s while getting card data", queryRowMessage)
	}

	var outData entity.CardData
	err = row.Scan(&outData.Name, &outData.Number, &outData.ExpirationMonth, &outData.ExpirationYear, &outData.Code, &outData.Cardholder, &outData.Metadata.Data)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.ErrLoginPasswordDataNotFound
		}

		return fmt.Errorf("error while processing response row in postgres while getting card data: %w", err)
	}

	query := `DELETE FROM card_data WHERE user_login=$1 AND name=$2`

	_, err = tx.Exec(ctx, query, userID, name)
	if err != nil {
		return fmt.Errorf("couldn't delete card data from postgres request: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("couldn't commit transaction while deleting card data postgres: %w", err)
	}

	return nil
}

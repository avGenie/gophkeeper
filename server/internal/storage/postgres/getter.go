package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/storage"
)

// GetLoginPasswordData Returns login-password data from postgres storage for user
func (p *Postgres) GetLoginPasswordData(ctx context.Context, name entity.DataName, userID entity.UserID) (entity.LoginPassword, error) {
	query := `SELECT name, login, password, meta FROM login_password_data WHERE user_login=$1 AND name=$2`

	row := p.pool.QueryRow(ctx, query, userID, name)
	if row == nil {
		return entity.LoginPassword{}, fmt.Errorf("%s while getting login-password data", queryRowMessage)
	}

	var outData entity.LoginPassword
	err := row.Scan(&outData.Name, &outData.Login, &outData.Password, &outData.Metadata.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.LoginPassword{}, storage.ErrLoginPasswordDataNotFound
		}

		return entity.LoginPassword{}, fmt.Errorf("error while processing response row in postgres while getting login-password data: %w", err)
	}

	return outData, nil
}

// GetTextData Returns text data from postgres storage for user
func (p *Postgres) GetTextData(ctx context.Context, name entity.DataName, userID entity.UserID) (entity.TextData, error) {
	query := `SELECT name, text, meta FROM text_data WHERE user_login=$1 AND name=$2`

	row := p.pool.QueryRow(ctx, query, userID, name)
	if row == nil {
		return entity.TextData{}, fmt.Errorf("%s while getting login-password data", queryRowMessage)
	}

	var outData entity.TextData
	err := row.Scan(&outData.Name, &outData.Text, &outData.Metadata.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.TextData{}, storage.ErrLoginPasswordDataNotFound
		}

		return entity.TextData{}, fmt.Errorf("error while processing response row in postgres while getting login-password data: %w", err)
	}

	return outData, nil
}

// GetCardData Returns card data from postgres storage for user
func (p *Postgres) GetCardData(ctx context.Context, name entity.DataName, userID entity.UserID) (entity.CardData, error) {
	query := `SELECT name, number, expiration_month, expiration_year, code, cardholder, meta FROM cards_data WHERE user_login=$1 AND name=$2`

	row := p.pool.QueryRow(ctx, query, userID, name)
	if row == nil {
		return entity.CardData{}, fmt.Errorf("%s while getting login-password data", queryRowMessage)
	}

	var outData entity.CardData
	err := row.Scan(&outData.Name, &outData.Number, &outData.ExpirationMonth, &outData.ExpirationYear, &outData.Code, &outData.Cardholder, &outData.Metadata.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.CardData{}, storage.ErrLoginPasswordDataNotFound
		}

		return entity.CardData{}, fmt.Errorf("error while processing response row in postgres while getting login-password data: %w", err)
	}

	return outData, nil
}

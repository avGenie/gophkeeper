package postgres

import (
	"context"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
)

// UpdateLoginPasswordData Updates login-password data to postgres storage
func (p *Postgres) UpdateLoginPasswordData(ctx context.Context, data entity.LoginPassword, userID entity.UserID) error {
	query := `UPDATE login_password_data SET name=$1, login=$2, password=$3, meta=$4 WHERE user_login=$5`

	_, err := p.pool.Exec(ctx, query, data.Name, data.Login, data.Password, data.Metadata.Data, userID)
	if err != nil {
		return fmt.Errorf("couldn't update login-password data postgres: %w", err)
	}

	return nil
}

// UpdateTextData Updates text data to postgres storage
func (p *Postgres) UpdateTextData(ctx context.Context, data entity.TextData, userID entity.UserID) error {
	query := `UPDATE text_data SET name=$1, text=$2, meta=$3 WHERE user_login=$4`

	_, err := p.pool.Exec(ctx, query, data.Name, data.Text, data.Metadata.Data, userID)
	if err != nil {
		return fmt.Errorf("couldn't update text data postgres: %w", err)
	}

	return nil
}

// UpdateCardData Updates card data to postgres storage
func (p *Postgres) UpdateCardData(ctx context.Context, data entity.CardData, userID entity.UserID) error {
	query := `UPDATE cards_data SET name=$1, number=$2, expiration_month=$3, expiration_year=$4, code=$5, cardholder=$6, meta=$7 WHERE user_login=$8`

	_, err := p.pool.Exec(ctx, query, data.Name, data.Number, data.ExpirationMonth, data.ExpirationYear, data.Code, data.Cardholder, data.Metadata.Data, userID)
	if err != nil {
		return fmt.Errorf("couldn't update card data postgres: %w", err)
	}

	return nil
}

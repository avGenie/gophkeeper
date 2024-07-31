package postgres

import (
	"context"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
)

// UpdateLoginPasswordData Updates login-password data to postgres storage
func (p *Postgres) UpdateLoginPasswordData(ctx context.Context, data entity.LoginPassword, userID entity.UserID) error {
	query := `UPDATE login_password_data SET login=$1, password=$2, meta=$3 WHERE user_login=$4 AND name=$5`

	_, err := p.pool.Exec(ctx, query, data.Login, data.Password, data.Metadata.Data, userID, data.Name)
	if err != nil {
		return fmt.Errorf("couldn't update login-password data postgres: %w", err)
	}

	return nil
}

// UpdateTextData Updates text data to postgres storage
func (p *Postgres) UpdateTextData(ctx context.Context, data entity.TextData, userID entity.UserID) error {
	query := `UPDATE text_data SET text=$1, meta=$2 WHERE user_login=$3 AND name=$4`

	_, err := p.pool.Exec(ctx, query, data.Text, data.Metadata.Data, userID, data.Name)
	if err != nil {
		return fmt.Errorf("couldn't update text data postgres: %w", err)
	}

	return nil
}

// UpdateCardData Updates card data to postgres storage
func (p *Postgres) UpdateCardData(ctx context.Context, data entity.CardData, userID entity.UserID) error {
	query := `UPDATE cards_data SET number=$1, expiration_month=$2, expiration_year=$3, code=$4, cardholder=$5, meta=$6 WHERE user_login=$7 AND name=$8`

	_, err := p.pool.Exec(ctx, query, data.Number, data.ExpirationMonth, data.ExpirationYear, data.Code, data.Cardholder, data.Metadata.Data, userID, data.Name)
	if err != nil {
		return fmt.Errorf("couldn't update card data postgres: %w", err)
	}

	return nil
}

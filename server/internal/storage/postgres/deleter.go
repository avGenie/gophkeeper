package postgres

import (
	"context"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
)

// DeleteLoginPasswordData Deletes login-password data from postgres storage for user
func (p *Postgres) DeleteLoginPasswordData(ctx context.Context, name entity.DataName, userID entity.UserID) error {
	query := `DELETE FROM login_password_data WHERE user_login=$1 AND name=$2`

	_, err := p.pool.Exec(ctx, query, userID, name)
	if err != nil {
		return fmt.Errorf("couldn't delete login-password data from postgres request: %w", err)
	}

	return nil
}

// DeleteTextData Deletes text data from postgres storage for user
func (p *Postgres) DeleteTextData(ctx context.Context, name entity.DataName, userID entity.UserID) error {
	query := `DELETE FROM text_data WHERE user_login=$1 AND name=$2`

	_, err := p.pool.Exec(ctx, query, userID, name)
	if err != nil {
		return fmt.Errorf("couldn't delete text data from postgres request: %w", err)
	}

	return nil
}

// DeleteCardData Deletes card data from postgres storage for user
func (p *Postgres) DeleteCardData(ctx context.Context, name entity.DataName, userID entity.UserID) error {
	query := `DELETE FROM card_data WHERE user_login=$1 AND name=$2`

	_, err := p.pool.Exec(ctx, query, userID, name)
	if err != nil {
		return fmt.Errorf("couldn't delete card data from postgres request: %w", err)
	}

	return nil
}

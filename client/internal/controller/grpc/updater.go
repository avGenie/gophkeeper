package grpc

import (
	"context"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	client_context "github.com/avGenie/gophkeeper/client/internal/usecase/context"
	"github.com/avGenie/gophkeeper/client/internal/usecase/converter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateLoginPasswordData Update login-password data for user
func (c *GRPCClient) UpdateLoginPasswordData(object entity.LoginPassword) error {
	dataRequest := converter.ConvertLoginPasswordDataToPbLoginPassword(object)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, c.userToken)

	_, err := c.client.UpdateLoginPassword(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			return fmt.Errorf("failed to update login-password objects: %w", err)
		}

		return fmt.Errorf("failed to update login-password data: %w", err)
	}

	return nil
}

// UpdateTextData Update text data for user
func (c *GRPCClient) UpdateTextData(object entity.TextData) error {
	dataRequest := converter.ConvertTextToPbText(object)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, c.userToken)

	_, err := c.client.UpdateText(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			return fmt.Errorf("failed to update text objects: %w", err)
		}

		return fmt.Errorf("failed to update text data: %w", err)
	}

	return nil
}

// UpdateCardData Update card data for user
func (c *GRPCClient) UpdateCardData(object entity.CardData) error {
	dataRequest := converter.ConvertCardToPbCard(object)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, c.userToken)

	_, err := c.client.UpdateCard(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			return fmt.Errorf("failed to update card objects: %w", err)
		}

		return fmt.Errorf("failed to update card data: %w", err)
	}

	return nil
}

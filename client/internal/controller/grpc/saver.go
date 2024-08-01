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

// SaveLoginPasswordData Save login-password data for user
func (c *GRPCClient) SaveLoginPasswordData(object entity.LoginPassword) error {
	dataRequest := converter.ConvertLoginPasswordDataToPbLoginPassword(object)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, c.userToken)

	_, err := c.client.SaveLoginPassword(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			if status.Code(err) == codes.AlreadyExists {
				return controller.ErrDataAlreadyExists
			}

			return fmt.Errorf("failed to save login-password objects: %w", err)
		}

		return fmt.Errorf("failed to save login-password data: %w", err)
	}

	return nil
}

// SaveTextData Save text data for user
func (c *GRPCClient) SaveTextData(object entity.TextData) error {
	dataRequest := converter.ConvertTextToPbText(object)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, c.userToken)

	_, err := c.client.SaveText(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			if status.Code(err) == codes.AlreadyExists {
				return controller.ErrDataAlreadyExists
			}

			return fmt.Errorf("failed to save text objects: %w", err)
		}

		return fmt.Errorf("failed to save text data: %w", err)
	}

	return nil
}

// SaveCardData Save card data for user
func (c *GRPCClient) SaveCardData(object entity.CardData) error {
	dataRequest := converter.ConvertCardToPbCard(object)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, c.userToken)

	_, err := c.client.SaveCard(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			if status.Code(err) == codes.AlreadyExists {
				return controller.ErrDataAlreadyExists
			}

			return fmt.Errorf("failed to save card objects: %w", err)
		}

		return fmt.Errorf("failed to save card data: %w", err)
	}

	return nil
}

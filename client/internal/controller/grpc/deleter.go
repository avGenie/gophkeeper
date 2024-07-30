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

// DeleteLoginPasswordData Delete login-password data for user
func (c *GRPCClient) DeleteLoginPasswordData(name entity.ObjectName, token entity.Token) error {
	dataRequest := converter.CreatePbDataGetterRequest(name, entity.DataRequestLoginPassword)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	_, err := c.client.DeleteObject(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			if status.Code(err) == codes.NotFound {
				return controller.ErrDataDeleteWrongName
			}

			return fmt.Errorf("failed to delete login-password objects: %w", err)
		}

		return fmt.Errorf("failed to delete login-password data: %w", err)
	}

	return nil
}

// DeleteTextData Delete text data for user
func (c *GRPCClient) DeleteTextData(name entity.ObjectName, token entity.Token) error {
	dataRequest := converter.CreatePbDataGetterRequest(name, entity.DataRequestText)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	_, err := c.client.DeleteObject(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			if status.Code(err) == codes.NotFound {
				return controller.ErrDataDeleteWrongName
			}

			return fmt.Errorf("failed to delete login-password objects: %w", err)
		}

		return fmt.Errorf("failed to delete login-password data: %w", err)
	}

	return nil
}

// DeleteCardData Delete card data for user
func (c *GRPCClient) DeleteCardData(name entity.ObjectName, token entity.Token) error {
	dataRequest := converter.CreatePbDataGetterRequest(name, entity.DataRequestCard)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	_, err := c.client.DeleteObject(ctx, dataRequest)
	if err != nil {
		if err != nil {
			if status.Code(err) == codes.PermissionDenied {
				return controller.ErrUserPermissionDenied
			}

			if status.Code(err) == codes.InvalidArgument {
				return controller.ErrDataDeleteWrongName
			}

			if status.Code(err) == codes.NotFound {
				return controller.ErrDataDeleteWrongName
			}

			return fmt.Errorf("failed to delete login-password objects: %w", err)
		}

		return fmt.Errorf("failed to delete login-password data: %w", err)
	}

	return nil
}

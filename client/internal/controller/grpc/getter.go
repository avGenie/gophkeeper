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
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetLoginPasswordData Returns login-password data for user
func (c *GRPCClient) GetLoginPasswordData(name entity.ObjectName, token entity.Token) (entity.LoginPassword, error) {
	dataRequest := converter.ConvertObjectNameToPbDataRequest(name)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	data, err := c.client.GetLoginPasswordObject(ctx, dataRequest)
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			return entity.LoginPassword{}, controller.ErrUserPermissionDenied
		}

		if status.Code(err) == codes.NotFound {
			return entity.LoginPassword{}, controller.ErrDataNotFound
		}

		return entity.LoginPassword{}, fmt.Errorf("failed to get login-password data: %w", err)
	}

	return converter.ConvertPbLoginPasswordDataToLoginPasswordData(data), nil
}

// GetLoginPasswordObjects Returns login-password objects for user
func (c *GRPCClient) GetLoginPasswordObjects(token entity.Token) (entity.LoginPasswordObjects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	obj, err := c.client.GetLoginPasswordObjects(ctx, &emptypb.Empty{})
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			return entity.LoginPasswordObjects{}, controller.ErrUserPermissionDenied
		}

		if status.Code(err) == codes.NotFound {
			return entity.LoginPasswordObjects{}, controller.ErrDataNotFound
		}

		return entity.LoginPasswordObjects{}, fmt.Errorf("failed to get login-password objects: %w", err)
	}

	return converter.ConvertPbLoginPasswordObjectsToLoginPasswordObjects(obj), nil
}

// GetTextUser Returns text data for user
func (c *GRPCClient) GetTextData(name entity.ObjectName, token entity.Token) (entity.TextData, error) {
	dataRequest := converter.ConvertObjectNameToPbDataRequest(name)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	data, err := c.client.GetTextObject(ctx, dataRequest)
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			return entity.TextData{}, controller.ErrUserPermissionDenied
		}

		if status.Code(err) == codes.NotFound {
			return entity.TextData{}, controller.ErrDataNotFound
		}

		return entity.TextData{}, fmt.Errorf("failed to get text data: %w", err)
	}

	return converter.ConvertPbTextToText(data), nil
}

// GetTextObjects Returns text objects for user
func (c *GRPCClient) GetTextObjects(token entity.Token) (entity.TextObjects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	obj, err := c.client.GetTextObjects(ctx, &emptypb.Empty{})
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			return entity.TextObjects{}, controller.ErrUserPermissionDenied
		}

		if status.Code(err) == codes.NotFound {
			return entity.TextObjects{}, controller.ErrDataNotFound
		}

		return entity.TextObjects{}, fmt.Errorf("failed to get login-password objects: %w", err)
	}

	return converter.ConvertPbTextObjectsToTextObject(obj), nil
}

// GetCardData Returns card data for user
func (c *GRPCClient) GetCardData(name entity.ObjectName, token entity.Token) (entity.CardData, error) {
	dataRequest := converter.ConvertObjectNameToPbDataRequest(name)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	data, err := c.client.GetCardObject(ctx, dataRequest)
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			return entity.CardData{}, controller.ErrUserPermissionDenied
		}

		if status.Code(err) == codes.NotFound {
			return entity.CardData{}, controller.ErrDataNotFound
		}

		return entity.CardData{}, fmt.Errorf("failed to get text data: %w", err)
	}

	return converter.ConvertPbCardToCard(data), nil
}

// GetCardObjects Returns card objects for user
func (c *GRPCClient) GetCardObjects(token entity.Token) (entity.CardObjects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	ctx = client_context.SaveTokenToContext(ctx, token)

	obj, err := c.client.GetCardObjects(ctx, &emptypb.Empty{})
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			return entity.CardObjects{}, controller.ErrUserPermissionDenied
		}

		if status.Code(err) == codes.NotFound {
			return entity.CardObjects{}, controller.ErrDataNotFound
		}

		return entity.CardObjects{}, fmt.Errorf("failed to get login-password objects: %w", err)
	}

	return converter.ConvertPbCardObjectsToCardObject(obj), nil
}

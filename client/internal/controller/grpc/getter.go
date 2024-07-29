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

func (c *GRPCClient) GetLoginPasswordUser(name entity.ObjectName, token entity.Token) (entity.LoginPassword, error) {
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

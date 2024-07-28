package grpc

import (
	"context"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"github.com/avGenie/gophkeeper/client/internal/usecase/converter"
	client_context "github.com/avGenie/gophkeeper/client/internal/usecase/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

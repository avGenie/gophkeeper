package grpc

import (
	"context"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"github.com/avGenie/gophkeeper/client/internal/usecase/converter"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterUser Implements user registration
func (c *GRPCClient) RegisterUser(userCreds entity.User) error {
	pbCreds := converter.ConvertUserCredentialsToPbUserCredentials(userCreds)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	_, err := c.client.RegisterUser(ctx, pbCreds)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return controller.ErrUserExists
		}

		if status.Code(err) == codes.InvalidArgument {
			return controller.ErrUserInvalid
		}

		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

// AuthenticateUser Implements user authentication. Returns token.
func (c *GRPCClient) AuthenticateUser(userCreds entity.User) error {
	pbCreds := converter.ConvertUserCredentialsToPbUserCredentials(userCreds)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	token, err := c.client.AuthenticateUser(ctx, pbCreds)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return controller.ErrUserNotFound
		}

		if status.Code(err) == codes.InvalidArgument {
			return controller.ErrUserInvalid
		}

		return fmt.Errorf("failed to authenticate user: %w", err)
	}

	c.userToken = converter.ConvertPbAuthTokenToToken(token)

	zap.S().Debug("token created", zap.String("token", string(c.userToken)))

	return nil
}

package grpc

import (
	"context"
	"errors"

	pb "github.com/avGenie/gophkeeper/proto"
	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/avGenie/gophkeeper/server/internal/usecase/converter"
	"github.com/avGenie/gophkeeper/server/internal/usecase/crypto"
	"github.com/avGenie/gophkeeper/server/internal/usecase/validation"
	"go.uber.org/zap"
)

func (s *GRPCServer) RegisterUser(ctx context.Context, userCreds *pb.UserCredentials) (*pb.ErrorMessage, error) {
	user := converter.ConvertPbUserCredentialsToUser(userCreds)

	if !validation.ValidateCreateUserRequest(user) {
		zap.S().Error(FailedUserCredentials)

		return createPbErrorMessage(FailedUserCredentials), nil
	}

	user.ID = crypto.CreateUserID()

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	err := s.storage.CreateUser(ctx, user)
	if err != nil {
		zap.S().Error("failed to create user", zap.Error(err))

		if errors.Is(err, storage.ErrUserExists) {
			return createPbErrorMessage(err.Error()), nil
		}

		return createPbErrorMessage(InternalServerError), nil
	}

	return createPbSuccessMessage(), nil
}

func (s *GRPCServer) AuthenticateUser(ctx context.Context, userCreds *pb.UserCredentials) (*pb.AuthenticateResponse, error) {
	user := converter.ConvertPbUserCredentialsToUser(userCreds)

	if !validation.ValidateCreateUserRequest(user) {
		zap.S().Error(FailedUserCredentials)

		return createPbAuthResponse(createPbErrorMessage(FailedUserCredentials), entity.Token("")), nil
	}

	return createPbAuthResponse(createPbSuccessMessage(), entity.Token("")), nil
}

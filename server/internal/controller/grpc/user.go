package grpc

import (
	"context"
	"errors"

	pb "github.com/avGenie/gophkeeper/proto"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/avGenie/gophkeeper/server/internal/usecase/converter"
	"github.com/avGenie/gophkeeper/server/internal/usecase/crypto"
	"github.com/avGenie/gophkeeper/server/internal/usecase/validation"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) RegisterUser(ctx context.Context, userCreds *pb.UserCredentials) (*emptypb.Empty, error) {
	user := converter.ConvertPbUserCredentialsToUser(userCreds)

	if !validation.ValidateCreateUserRequest(user) {
		zap.S().Error(FailedUserCredentials)

		return nil, status.Errorf(codes.InvalidArgument, FailedUserCredentials)
	}

	user.ID = crypto.CreateUserID()

	user, err := crypto.HashPassword(user)
	if err != nil {
		zap.S().Error("failed to hash user password", zap.Error(err))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			zap.S().Info("failed to create user", zap.Error(err))

			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		zap.S().Error("failed to create user", zap.Error(err))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) AuthenticateUser(ctx context.Context, userCreds *pb.UserCredentials) (*pb.AuthorizationToken, error) {
	user := converter.ConvertPbUserCredentialsToUser(userCreds)

	if !validation.ValidateCreateUserRequest(user) {
		zap.S().Error(FailedUserCredentials)

		return nil, status.Errorf(codes.InvalidArgument, FailedUserCredentials)
	}

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	userStorage, err := s.storage.GetUser(ctx, user)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			zap.S().Info("failed to get user from storage", zap.Error(err))

			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		zap.S().Error("failed to get user from storage", zap.Error(err))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	err = crypto.CheckPasswordHash(user.Password, userStorage.Password)
	if err != nil {
		if errors.Is(err, crypto.ErrWrongPassword) {
			zap.S().Info("failed to check user password", zap.Error(err))

			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		zap.S().Error("failed to check user password", zap.Error(err))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	token, err := crypto.BuildJWTString(user.ID)
	if err != nil {
		zap.S().Error("failed to create token", zap.Error(err))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	return converter.ConvertTokenToPbAuthToken(token), nil
}

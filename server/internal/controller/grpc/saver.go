package grpc

import (
	"context"
	"errors"

	pb "github.com/avGenie/gophkeeper/proto"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/avGenie/gophkeeper/server/internal/usecase"
	"github.com/avGenie/gophkeeper/server/internal/usecase/converter"
	"github.com/avGenie/gophkeeper/server/internal/usecase/crypto"
	"github.com/avGenie/gophkeeper/server/internal/usecase/validation"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SaveLoginPassword Saves logn-password data for user
func (s *GRPCServer) SaveLoginPassword(ctx context.Context, loginPasswordData *pb.LoginPasswordData) (*emptypb.Empty, error) {
	userID, err := usecase.GetUserIDFromContext(ctx)
	if err != nil {
		zap.S().Error("error while getting user id from context", zap.Error(err))

		if errors.Is(err, usecase.ErrEmptyMetadata) || errors.Is(err, usecase.ErrEmptyTokenValue) ||
			errors.Is(err, crypto.ErrTokenNotValid) || errors.Is(err, crypto.ErrTokenExpired) {
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	loginPass := converter.ConvertPbLoginPasswordToLoginPassword(loginPasswordData)

	if !validation.ValidateLoginPasswordData(loginPass) {
		zap.S().Error("invalid login-password data", zap.String("user_id", string(userID)))

		return nil, status.Errorf(codes.InvalidArgument, FailedData)
	}

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	err = s.storage.SaveLoginPasswordData(ctx, loginPass, userID)
	if err != nil {
		if errors.Is(err, storage.ErrDataExists) {
			zap.S().Info("failed to save login-password data", zap.Error(err))

			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		zap.S().Error("failed to save login-password data", zap.Error(err))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	return &emptypb.Empty{}, nil
}

// SaveLoginPassword Saves logn-password data for user
func (s *GRPCServer) SaveText(ctx context.Context, textData *pb.TextData) (*emptypb.Empty, error) {
	userID, err := usecase.GetUserIDFromContext(ctx)
	if err != nil {
		zap.S().Error("error while getting user id from context", zap.Error(err))

		if errors.Is(err, usecase.ErrEmptyMetadata) || errors.Is(err, usecase.ErrEmptyTokenValue) ||
			errors.Is(err, crypto.ErrTokenNotValid) || errors.Is(err, crypto.ErrTokenExpired) {
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	data := converter.ConvertPbTextToText(textData)

	if !validation.ValidateTextData(data) {
		zap.S().Error("invalid login-password data", zap.String("user_id", string(userID)))

		return nil, status.Errorf(codes.InvalidArgument, FailedData)
	}

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	err = s.storage.SaveTextData(ctx, data, userID)
	if err != nil {
		if errors.Is(err, storage.ErrDataExists) {
			zap.S().Info("failed to save login-password data", zap.Error(err))

			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		zap.S().Error("failed to save login-password data", zap.Error(err))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	return &emptypb.Empty{}, nil
}

package grpc

import (
	"context"
	"errors"

	pb "github.com/avGenie/gophkeeper/proto"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/avGenie/gophkeeper/server/internal/usecase"
	"github.com/avGenie/gophkeeper/server/internal/usecase/converter"
	"github.com/avGenie/gophkeeper/server/internal/usecase/crypto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetLoginPasswordObject Returns logn-password data for user
func (s *GRPCServer) GetLoginPasswordObject(ctx context.Context, inputData *pb.DataRequest) (*pb.LoginPasswordData, error) {
	userID, err := usecase.GetUserIDFromContext(ctx)
	if err != nil {
		zap.S().Error("error while getting user id from context", zap.Error(err))

		if errors.Is(err, usecase.ErrEmptyMetadata) || errors.Is(err, usecase.ErrEmptyTokenValue) ||
			errors.Is(err, crypto.ErrTokenNotValid) || errors.Is(err, crypto.ErrTokenExpired) {
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	dataName := converter.ConvertPbDataRequestToDataName(inputData)

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	data, err := s.storage.GetLoginPasswordData(ctx, dataName, userID)
	if err != nil {
		if errors.Is(err, storage.ErrLoginPasswordDataNotFound) {
			zap.S().Info("failed to get login-password for user from storage", zap.Error(err), zap.String("user_id", string(userID)))

			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		zap.S().Error("failed to get user from storage", zap.Error(err), zap.String("user_id", string(userID)))

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	outputData := converter.ConvertLoginPasswordDataToPbLoginPassword(data)

	return outputData, nil
}

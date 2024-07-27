package grpc

import (
	"context"
	"errors"

	pb "github.com/avGenie/gophkeeper/proto"
	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/usecase"
	"github.com/avGenie/gophkeeper/server/internal/usecase/converter"
	"github.com/avGenie/gophkeeper/server/internal/usecase/crypto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteObject Deletes object for user
func (s *GRPCServer) DeleteObject(ctx context.Context, pbRequest *pb.DataGetterRequest) (*emptypb.Empty, error) {
	userID, err := usecase.GetUserIDFromContext(ctx)
	if err != nil {
		zap.S().Error("error while getting user id from context", zap.Error(err))

		if errors.Is(err, usecase.ErrEmptyMetadata) || errors.Is(err, usecase.ErrEmptyTokenValue) ||
			errors.Is(err, crypto.ErrTokenNotValid) || errors.Is(err, crypto.ErrTokenExpired) {
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	request := converter.ConvertPbDataGetterRequestToDataRequest(pbRequest)

	if request.Type == entity.DataRequestInvalid {
		zap.S().Error("invalid data request type", zap.String("user_id", string(userID)))

		return nil, status.Errorf(codes.InvalidArgument, "invalid data reques type")
	}

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	switch request.Type {
	case entity.DataRequestLoginPassword:
		err = s.storage.DeleteLoginPasswordData(ctx, request.Name, userID)
		if err != nil {
			zap.S().Error("failed to delete login password data", zap.Error(err), zap.String("user_id", string(userID)))

			return nil, status.Errorf(codes.Internal, InternalServerError)
		}
	case entity.DataRequestText:
		err = s.storage.DeleteTextData(ctx, request.Name, userID)
		if err != nil {
			zap.S().Error("failed to delete text data", zap.Error(err), zap.String("user_id", string(userID)))

			return nil, status.Errorf(codes.Internal, InternalServerError)
		}
	case entity.DataRequestCard:
		err = s.storage.DeleteCardData(ctx, request.Name, userID)
		if err != nil {
			zap.S().Error("failed to delete card data", zap.Error(err), zap.String("user_id", string(userID)))

			return nil, status.Errorf(codes.Internal, InternalServerError)
		}
	}

	return &emptypb.Empty{}, nil
}

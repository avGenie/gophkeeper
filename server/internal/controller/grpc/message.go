package grpc

import (
	pb "github.com/avGenie/gophkeeper/proto"
	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/usecase/converter"
)

func createPbErrorMessage(message string) *pb.ErrorMessage {
	return &pb.ErrorMessage{
		Status: pb.ErrorStatus_Error_Failed,
		Error:  message,
	}
}

func createPbSuccessMessage() *pb.ErrorMessage {
	return &pb.ErrorMessage{
		Status: pb.ErrorStatus_Error_Success,
	}
}

func createPbAuthResponse(errMsg *pb.ErrorMessage, token entity.Token) *pb.AuthenticateResponse {
	return &pb.AuthenticateResponse{
		Token: converter.ConvertTokenToPbAuthToken(token),
		Error: errMsg,
	}
}

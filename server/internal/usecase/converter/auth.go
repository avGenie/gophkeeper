package converter

import (
	"github.com/avGenie/gophkeeper/server/internal/entity"

	pb "github.com/avGenie/gophkeeper/proto"
)

func ConvertTokenToPbAuthToken(token entity.Token) *pb.AuthorizationToken {
	return &pb.AuthorizationToken{
		Token: string(token),
	}
}

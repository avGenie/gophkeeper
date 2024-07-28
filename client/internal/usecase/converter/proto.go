package converter

import (
	"github.com/avGenie/gophkeeper/client/internal/entity"
	pb "github.com/avGenie/gophkeeper/proto"
)

func ConvertUserCredentialsToPbUserCredentials(data entity.User) *pb.UserCredentials {
	return &pb.UserCredentials{
		Login: data.Login,
		Password: data.Password,
	}
}

func ConvertPbAuthTokenToToken(data *pb.AuthorizationToken) entity.Token {
	return entity.Token(data.Token)
}

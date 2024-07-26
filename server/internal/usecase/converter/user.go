package converter

import (
	"github.com/avGenie/gophkeeper/server/internal/entity"

	pb "github.com/avGenie/gophkeeper/proto"
)

// ConvertTokenToPbAuthToken Converts protobuf user creds app user creds
func ConvertPbUserCredentialsToUser(pbUser *pb.UserCredentials) entity.User {
	return entity.User{
		Login:    pbUser.GetLogin(),
		Password: pbUser.GetPassword(),
	}
}

package converter

import (
	"github.com/avGenie/gophkeeper/server/internal/entity"

	pb "github.com/avGenie/gophkeeper/proto"
)

// ConvertPbMetadataToMetadata Converts protobuf metadata to app metadata
func ConvertPbMetadataToMetadata(data *pb.MetaData) entity.Metadata {
	return entity.Metadata{
		Data: data.GetData(),
	}
}

// ConvertTokenToPbAuthToken Converts protobuf user creds to app user creds
func ConvertPbUserCredentialsToUser(pbUser *pb.UserCredentials) entity.User {
	return entity.User{
		Login:    pbUser.GetLogin(),
		Password: pbUser.GetPassword(),
	}
}

// ConvertTokenToPbAuthToken Converts app token to protobuf token
func ConvertTokenToPbAuthToken(token entity.Token) *pb.AuthorizationToken {
	return &pb.AuthorizationToken{
		Token: string(token),
	}
}

// ConvertPbLoginPasswordToLoginPassword Converts protobuf login-password data to app login-password data
func ConvertPbLoginPasswordToLoginPassword(data *pb.LoginPasswordData) entity.LoginPassword {
	return entity.LoginPassword{
		Name:     data.GetName(),
		Login:    data.GetLogin(),
		Password: data.GetPassword(),
		Metadata: ConvertPbMetadataToMetadata(data.GetMeta()),
	}
}

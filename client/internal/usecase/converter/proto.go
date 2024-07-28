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

func ConvertObjectNameToPbDataRequest(data entity.ObjectName) *pb.DataRequest {
	return &pb.DataRequest{
		Name: string(data),
	}
}

func ConvertPbMetadataToMetadata(data *pb.MetaData) entity.Metadata {
	return entity.Metadata{
		Data: data.GetData(),
	}
}

func ConvertPbLoginPasswordDataToLoginPasswordData(data *pb.LoginPasswordData) entity.LoginPassword {
	return entity.LoginPassword{
		Name: data.GetName(),
		Login: data.GetLogin(),
		Password: data.GetPassword(),
		Metadata: ConvertPbMetadataToMetadata(data.GetMeta()),
	}
}

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

// ConvertMetadataToMetadataPb Converts app metadata to protobuf metadata
func ConvertMetadataToMetadataPb(data entity.Metadata) *pb.MetaData {
	return &pb.MetaData{
		Data: data.Data,
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

// ConvertLoginPasswordDataToPbLoginPassword Converts app login-password data to protobuf login-password data
func ConvertLoginPasswordDataToPbLoginPassword(data entity.LoginPassword) *pb.LoginPasswordData {
	return &pb.LoginPasswordData{
		Name:     data.Name,
		Login:    data.Login,
		Password: data.Password,
		Meta:     ConvertMetadataToMetadataPb(data.Metadata),
	}
}

// ConvertPbTextToText Converts protobuf text data to app text data
func ConvertPbTextToText(data *pb.TextData) entity.TextData {
	return entity.TextData{
		Name:     data.GetName(),
		Text:     data.GetText(),
		Metadata: ConvertPbMetadataToMetadata(data.GetMeta()),
	}
}

// ConvertTextToPbText Converts app text data to protobuf text data
func ConvertTextToPbText(data entity.TextData) *pb.TextData {
	return &pb.TextData{
		Name: data.Name,
		Text: data.Text,
		Meta: ConvertMetadataToMetadataPb(data.Metadata),
	}
}

// ConvertPbCardToCard Converts protobuf card data to app card data
func ConvertPbCardToCard(data *pb.CardData) entity.CardData {
	return entity.CardData{
		Name:            data.GetName(),
		Number:          data.GetNumber(),
		ExpirationMonth: int(data.GetExpirationMonth()),
		ExpirationYear:  int(data.GetExpirationYear()),
		Code:            int(data.GetCode()),
		Cardholder:      data.GetCardholder(),
		Metadata:        ConvertPbMetadataToMetadata(data.GetMeta()),
	}
}

// ConvertCardToPbCard Converts app card data to protobuf card data
func ConvertCardToPbCard(data entity.CardData) *pb.CardData {
	return &pb.CardData{
		Name:            data.Name,
		Number:          data.Number,
		ExpirationMonth: int32(data.ExpirationMonth),
		ExpirationYear:  int32(data.ExpirationYear),
		Code:            int32(data.Code),
		Cardholder:      data.Cardholder,
		Meta:            ConvertMetadataToMetadataPb(data.Metadata),
	}
}

func ConvertPbDataRequestToDataName(data *pb.DataRequest) entity.DataName {
	return entity.DataName(data.GetName())
}

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

// ConvertLoginPasswordObjectsToPbLoginPasswordObject Converts app login-password objects to protobuf login-password objects
func ConvertLoginPasswordObjectsToPbLoginPasswordObject(data entity.LoginPasswordObjs) *pb.LoginPasswordObjects {
	output := &pb.LoginPasswordObjects{
		Objects: make([]*pb.LoginPasswordData, 0, len(data)),
	}

	for _, obj := range data {
		outputObj := ConvertLoginPasswordDataToPbLoginPassword(obj)

		output.Objects = append(output.Objects, outputObj)
	}
	return output
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

// ConvertTextObjectsToPbTextObject Converts app text objects to protobuf text objects
func ConvertTextObjectsToPbTextObject(data entity.TextDataObjects) *pb.TextObjects {
	output := &pb.TextObjects{
		Objects: make([]*pb.TextData, 0, len(data)),
	}

	for _, obj := range data {
		outputObj := ConvertTextToPbText(obj)

		output.Objects = append(output.Objects, outputObj)
	}
	return output
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

// ConvertCardObjectsToPbCardObject Converts app card objects to protobuf card objects
func ConvertCardObjectsToPbCardObject(data entity.CardDataObjects) *pb.CardObjects {
	output := &pb.CardObjects{
		Objects: make([]*pb.CardData, 0, len(data)),
	}

	for _, obj := range data {
		outputObj := ConvertCardToPbCard(obj)

		output.Objects = append(output.Objects, outputObj)
	}
	return output
}

// ConvertPbDataRequestToDataName Converts protobuf data request to app data name
func ConvertPbDataRequestToDataName(data *pb.DataRequest) entity.DataName {
	return entity.DataName(data.GetName())
}

// ConvertPbDataGetterRequestToDataRequest Converts protobuf data getter request to app data request
func ConvertPbDataGetterRequestToDataRequest(data *pb.DataGetterRequest) entity.DataRequest {
	return entity.DataRequest{
		Type: ConvertPbDataTypeToDataRequestType(data.Type),
		Name: entity.DataName(data.GetName()),
	}
}

// ConvertPbDataTypeToDataRequestType Converts protobuf data type to app data request type
func ConvertPbDataTypeToDataRequestType(data pb.DataType) entity.DataRequestType {
	switch data.Type() {
	case pb.DataType_Type_LoginPassword.Type():
		return entity.DataRequestLoginPassword
	case pb.DataType_Type_Text.Type():
		return entity.DataRequestText
	case pb.DataType_Type_Card.Type():
		return entity.DataRequestCard
	default:
		return entity.DataRequestInvalid
	}
}

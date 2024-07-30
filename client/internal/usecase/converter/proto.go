package converter

import (
	"github.com/avGenie/gophkeeper/client/internal/entity"
	pb "github.com/avGenie/gophkeeper/proto"
)

// ConvertUserCredentialsToPbUserCredentials Converts protobuf user data to app user credentials
func ConvertUserCredentialsToPbUserCredentials(data entity.User) *pb.UserCredentials {
	return &pb.UserCredentials{
		Login:    data.Login,
		Password: data.Password,
	}
}

// ConvertPbAuthTokenToToken Converts protobuf auth token to app user token
func ConvertPbAuthTokenToToken(data *pb.AuthorizationToken) entity.Token {
	return entity.Token(data.Token)
}

// ConvertObjectNameToPbDataRequest Converts protobuf object name to app data request
func ConvertObjectNameToPbDataRequest(data entity.ObjectName) *pb.DataRequest {
	return &pb.DataRequest{
		Name: string(data),
	}
}

// ConvertPbMetadataToMetadata Converts protobuf metadata to app metadata
func ConvertPbMetadataToMetadata(data *pb.MetaData) entity.Metadata {
	return entity.Metadata{
		Data: data.GetData(),
	}
}

// ConvertPbLoginPasswordDataToLoginPasswordData Converts protobuf login-password data to app login-password data
func ConvertPbLoginPasswordDataToLoginPasswordData(data *pb.LoginPasswordData) entity.LoginPassword {
	return entity.LoginPassword{
		Name:     data.GetName(),
		Login:    data.GetLogin(),
		Password: data.GetPassword(),
		Metadata: ConvertPbMetadataToMetadata(data.GetMeta()),
	}
}

// ConvertPbLoginPasswordObjectsToLoginPasswordObjects Converts protobuf login-password objects to app login-password objects
func ConvertPbLoginPasswordObjectsToLoginPasswordObjects(data *pb.LoginPasswordObjects) entity.LoginPasswordObjects {
	var output entity.LoginPasswordObjects

	for _, obj := range data.Objects {
		outputObj := ConvertPbLoginPasswordDataToLoginPasswordData(obj)

		output = append(output, outputObj)
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

// ConvertPbTextObjectsToTextObject Converts protobuf text objects to app text objects
func ConvertPbTextObjectsToTextObject(data *pb.TextObjects) entity.TextObjects {
	var output entity.TextObjects

	for _, obj := range data.Objects {
		outputObj := ConvertPbTextToText(obj)

		output = append(output, outputObj)
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

// ConvertPbCardObjectsToCardObject Converts protobuf card objects to app card objects
func ConvertPbCardObjectsToCardObject(data *pb.CardObjects) entity.CardObjects {
	var output entity.CardObjects

	for _, obj := range data.Objects {
		outputObj := ConvertPbCardToCard(obj)

		output = append(output, outputObj)
	}

	return output
}

// ConvertDataRequestTypeToPbDataType Converts app data request object to protobuf data request object
func ConvertDataRequestTypeToPbDataType(requestType entity.DataRequestType) pb.DataType {
	switch requestType {
	case entity.DataRequestLoginPassword:
		return pb.DataType_Type_LoginPassword
	case entity.DataRequestText:
		return pb.DataType_Type_Text
	case entity.DataRequestCard:
		return pb.DataType_Type_Card
	default:
		return pb.DataType_Type_Invalid
	}
}

// CreatePbDataGetterRequest Creates protobuf data getter request object
func CreatePbDataGetterRequest(name entity.ObjectName, requestType entity.DataRequestType) *pb.DataGetterRequest {
	return &pb.DataGetterRequest{
		Type: ConvertDataRequestTypeToPbDataType(requestType),
		Name: string(name),
	}
}

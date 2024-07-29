package converter

import (
	pb "github.com/avGenie/gophkeeper/proto"
	"github.com/avGenie/gophkeeper/test-server-client/internal/entity"
)

func CreatePbMetadata(meta string) *pb.MetaData {
	return &pb.MetaData{
		Data: meta,
	}
}

func CreatePbDataRequest(name string) *pb.DataRequest {
	return &pb.DataRequest{
		Name: name,
	}
}

func CreatePbLoginPasswordData(name, login, password, meta string) *pb.LoginPasswordData {
	return &pb.LoginPasswordData{
		Name:     name,
		Login:    login,
		Password: password,
		Meta:     CreatePbMetadata(meta),
	}
}

func CreatePbTextData(name, text, meta string) *pb.TextData {
	return &pb.TextData{
		Name: name,
		Text: text,
		Meta: CreatePbMetadata(meta),
	}
}

func CreatePbCardData(name, number, cardholder string, expireMonth, expireYear, code int, meta string) *pb.CardData {
	return &pb.CardData{
		Name:            name,
		Number:          number,
		ExpirationMonth: int32(expireMonth),
		ExpirationYear:  int32(expireYear),
		Code:            int32(code),
		Cardholder:      cardholder,
		Meta:            CreatePbMetadata(meta),
	}
}


func CreatePbDataGetterRequest(name string, requestType entity.DataRequestType) *pb.DataGetterRequest {
	return &pb.DataGetterRequest{
		Type: ConvertDataRequestTypeToPbDataType(requestType),
		Name: name,
	}
}

func ConvertPbLoginPasswordDataToLoginPassword(data *pb.LoginPasswordData) entity.LoginPassword {
	return entity.LoginPassword{
		Name:     data.GetName(),
		Login:    data.GetLogin(),
		Password: data.GetPassword(),
		Metadata: data.GetMeta().Data,
	}
}

func ConvertPbLoginPasswordObjectsToLoginPasswordObjects(data *pb.LoginPasswordObjects) entity.LoginPasswordObjects {
	var output entity.LoginPasswordObjects

	for _, obj := range data.Objects {
		outputObj := ConvertPbLoginPasswordDataToLoginPassword(obj)

		output = append(output, outputObj)
	}

	return output
}

func ConvertPbTextDataToText(data *pb.TextData) entity.TextData {
	return entity.TextData{
		Name:     data.GetName(),
		Text:     data.GetText(),
		Metadata: data.GetMeta().Data,
	}
}

func ConvertPbTextObjectsToTextObjects(data *pb.TextObjects) entity.TextObjects {
	var output entity.TextObjects

	for _, obj := range data.Objects {
		outputObj := ConvertPbTextDataToText(obj)

		output = append(output, outputObj)
	}

	return output
}

func ConvertPbCardDataToCard(data *pb.CardData) entity.CardData {
	return entity.CardData{
		Name:            data.GetName(),
		Number:          data.GetNumber(),
		ExpirationMonth: int(data.GetExpirationMonth()),
		ExpirationYear:  int(data.GetExpirationYear()),
		Code:            int(data.GetCode()),
		Cardholder:      data.GetCardholder(),
		Metadata:        data.GetMeta().Data,
	}
}

func ConvertPbCardObjectsToCardObjects(data *pb.CardObjects) entity.CardObjects {
	var output entity.CardObjects

	for _, obj := range data.Objects {
		outputObj := ConvertPbCardDataToCard(obj)

		output = append(output, outputObj)
	}

	return output
}

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

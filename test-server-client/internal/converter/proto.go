package converter

import (
	pb "github.com/avGenie/gophkeeper/proto"
)

func CreatePbMetadata(meta string) *pb.MetaData {
	return &pb.MetaData{
		Data: meta,
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

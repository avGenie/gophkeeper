package loginpassword

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

type LoginPasswordProcessor struct {
	client controller.GophkeeperClient

	data   map[entity.ObjectName]entity.LoginPassword
}

func NewLoginPassword(client controller.GophkeeperClient) *LoginPasswordProcessor {
	return &LoginPasswordProcessor{
		client: client,
		data: make(map[entity.ObjectName]entity.LoginPassword),
	}
}

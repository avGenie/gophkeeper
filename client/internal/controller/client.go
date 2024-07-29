package controller

import "github.com/avGenie/gophkeeper/client/internal/entity"

type GophkeeperClient interface {
	Start()
	Stop()

	RegisterUser(userCreds entity.User) error
	AuthenticateUser(userCreds entity.User) (entity.Token, error)

	GetLoginPasswordData(name entity.ObjectName, token entity.Token) (entity.LoginPassword, error)
	GetLoginPasswordObjects(token entity.Token) (entity.LoginPasswordObjects, error)
	GetTextData(name entity.ObjectName, token entity.Token) (entity.TextData, error)
	GetTextObjects(token entity.Token) (entity.TextObjects, error)
}

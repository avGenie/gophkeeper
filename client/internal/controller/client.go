package controller

import "github.com/avGenie/gophkeeper/client/internal/entity"

type GophkeeperClient interface {
	Start()
	Stop()

	RegisterUser(userCreds entity.User) error
	AuthenticateUser(userCreds entity.User) (entity.Token, error)

	GetLoginPasswordUser(name entity.ObjectName, token entity.Token) (entity.LoginPassword, error)
}

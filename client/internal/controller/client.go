package controller

import "github.com/avGenie/gophkeeper/client/internal/entity"

type GophkeeperClient interface {
	Start()
	Stop()

	RegisterUser(userCreds entity.User) error
	AuthenticateUser(userCreds entity.User) (entity.Token, error)

	GetLoginPasswordData(name entity.ObjectName) (entity.LoginPassword, error)
	GetLoginPasswordObjects() (entity.LoginPasswordObjects, error)
	GetTextData(name entity.ObjectName) (entity.TextData, error)
	GetTextObjects() (entity.TextObjects, error)
	GetCardData(name entity.ObjectName) (entity.CardData, error)
	GetCardObjects() (entity.CardObjects, error)

	DeleteLoginPasswordData(name entity.ObjectName, token entity.Token) error
	DeleteTextData(name entity.ObjectName, token entity.Token) error
	DeleteCardData(name entity.ObjectName, token entity.Token) error
}

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
	GetCardData(name entity.ObjectName, token entity.Token) (entity.CardData, error)
	GetCardObjects(token entity.Token) (entity.CardObjects, error)

	DeleteLoginPasswordData(name entity.ObjectName, token entity.Token) error
	DeleteTextData(name entity.ObjectName, token entity.Token) error
	DeleteCardData(name entity.ObjectName, token entity.Token) error
}

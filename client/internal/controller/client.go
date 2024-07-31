package controller

import "github.com/avGenie/gophkeeper/client/internal/entity"

type GophkeeperClient interface {
	Start()
	Stop()

	RegisterUser(userCreds entity.User) error
	AuthenticateUser(userCreds entity.User) error

	GetLoginPasswordData(name entity.ObjectName) (entity.LoginPassword, error)
	GetLoginPasswordObjects() (entity.LoginPasswordObjects, error)
	GetTextData(name entity.ObjectName) (entity.TextData, error)
	GetTextObjects() (entity.TextObjects, error)
	GetCardData(name entity.ObjectName) (entity.CardData, error)
	GetCardObjects() (entity.CardObjects, error)

	DeleteLoginPasswordData(name entity.ObjectName) error
	DeleteTextData(name entity.ObjectName) error
	DeleteCardData(name entity.ObjectName) error

	SaveLoginPasswordData(object entity.LoginPassword) error
	SaveTextData(object entity.TextData) error
	SaveCardData(object entity.CardData) error
}

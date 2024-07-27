package storage

import (
	"context"

	"github.com/avGenie/gophkeeper/server/internal/entity"
)

// Storage Interface for storage implementation
type Storage interface {
	Close()

	CreateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, user entity.User) (entity.User, error)

	SaveLoginPasswordData(ctx context.Context, data entity.LoginPassword, userID entity.UserID) error
	SaveTextData(ctx context.Context, data entity.TextData, userID entity.UserID) error
	SaveCardData(ctx context.Context, data entity.CardData, userID entity.UserID) error

	GetLoginPasswordData(ctx context.Context, name entity.DataName, userID entity.UserID) (entity.LoginPassword, error)
	GetTextData(ctx context.Context, name entity.DataName, userID entity.UserID) (entity.TextData, error)
	GetCardData(ctx context.Context, name entity.DataName, userID entity.UserID) (entity.CardData, error)

	DeleteLoginPasswordData(ctx context.Context, name entity.DataName, userID entity.UserID) error
	DeleteTextData(ctx context.Context, name entity.DataName, userID entity.UserID) error
	DeleteCardData(ctx context.Context, name entity.DataName, userID entity.UserID) error
}

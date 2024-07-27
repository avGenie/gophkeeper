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
}

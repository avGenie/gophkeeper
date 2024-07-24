package storage

import (
	"context"

	"github.com/avGenie/gophkeeper/server/internal/app/entity"
)

// Storage Interface for storage implementation
type Storage interface {
	CreateUser(ctx context.Context, user entity.User) error
	AuthorizeUser(ctx context.Context, user entity.User) error
}

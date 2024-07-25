package crypto

import (
	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/google/uuid"
)

func CreateUserID() entity.UserID {
	uuid := uuid.New()

	return entity.UserID(uuid.String())
}

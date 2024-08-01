package crypto

import (
	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/google/uuid"
)

// CreateUUID Creates uuid using google algorithm
func CreateUUID() entity.UserID {
	uuid := uuid.New()

	return entity.UserID(uuid.String())
}

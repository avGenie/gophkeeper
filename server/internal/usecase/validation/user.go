package validation

import "github.com/avGenie/gophkeeper/server/internal/entity"

func ValidateCreateUserRequest(user entity.User) bool {
	return len(user.Login) > 0 && len(user.Password) > 0
}